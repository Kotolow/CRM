using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using CRMAuth.Models;
using CRMAuth.Data;
using DotNetEnv;
using Microsoft.IdentityModel.Tokens;

namespace CRMAuth.Services;

public class UserService : IUserService
{
    private readonly IUserRepository _userRepository;

    public UserService(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }
    public async Task<ApiResponse<int>> CreateUserAsync(CreateUserDto model)
    {
        var user = new User
        {
            Id = 0,
            Name = model.name,
            Email = model.email,
            PasswordHash = HashPassword(model.password),
            AvatarUrl = model.avatar_url,
            CreatedAt = DateTime.UtcNow,
            UpdatedAt = null
        };

        await _userRepository.AddUserAsync(user);

        return new ApiResponse<int>
        {
            status_code = 201,
            message = "User created successfully",
            data = user.Id
        };
    }

    public async Task<RegistrationDto> GetRegInfo(int id)
    {
        var user = await _userRepository.GetUserByIdAsync(id);
        return new RegistrationDto
        {
            id = user.Id,
            name = user.Name,
            avatar_url = user.AvatarUrl,
            created_at = user.CreatedAt,
            email = user.Email
        };
    }
    
    public async Task UpdateJwtTokenAsync(int userId, string token)
    {
        await _userRepository.UpdateJwtTokenAsync(userId, token);
    }

    public async Task<LoginResponseDto> LoginAsync(LoginDto loginDto)
    {
        var user = await GetUserByEmail(loginDto.Email);
        if (user == null || !VerifyPassword(loginDto.Password, user.PasswordHash))
        {
            return null;
        }

        var token = GenerateJwtToken(user);
        await _userRepository.UpdateJwtTokenAsync(user.Id, token);

        return new LoginResponseDto { JwtToken = token };
    }

    public async Task<User?> GetUserInfoAsync(int id)
    {
        var user = await _userRepository.GetUserByIdAsync(id);
        if (user == null)
        {
            return null;
        }

        return user;
    }
    
    public async Task<ApiResponse<UserDto?>> GetUserByIdAsync(int id)
    {
        var user = await _userRepository.GetUserByIdAsync(id);
        if (user == null)
        {
            return new ApiResponse<UserDto?>
            {
                status_code = 404,
                message = "User not found",
                data = null
            };
        }

        return new ApiResponse<UserDto?>
        {
            status_code = 200,
            message = "Success",
            data = new UserDto { name = user.Name, email = user.Email, avatar_url = user.AvatarUrl }
        };
    }
    
    public async Task<User?> GetUserByEmail(string email)
    {
        var users = await _userRepository.GetAllUsersAsync();
        return users.FirstOrDefault(x => x.Email == email);
    }
    
    public async Task<ApiResponse<bool>> UpdateUserAsync(int id, UpdateUserDto model)
    {
        var user = await _userRepository.GetUserByIdAsync(id);
        if (user == null)
        {
            return new ApiResponse<bool>
            {
                status_code = 404,
                message = "User not found",
                data = false
            };
        }

        user.Name = model.name ?? user.Name;
        user.AvatarUrl = model.avatar_url ?? user.AvatarUrl;
        user.UpdatedAt = DateTime.UtcNow;

        await _userRepository.UpdateUserAsync(user);

        return new ApiResponse<bool>
        {
            status_code = 200,
            message = "User updated successfully",
            data = true
        };
    }

    public async Task<ApiResponse<List<UserDto>>> GetAllUsersAsync()
    {
        var users = await _userRepository.GetAllUsersAsync();
        var userDtos = users.Select(user => new UserDto
        {
            name = user.Name,
            email = user.Email,
            avatar_url = user.AvatarUrl
        }).ToList();

        return new ApiResponse<List<UserDto>>
        {
            status_code = 200,
            message = "Success",
            data = userDtos
        };
    }

    public string GenerateJwtToken(User user)
    {
        Env.Load();
        var tokenHandler = new JwtSecurityTokenHandler();
        var key = Encoding.ASCII.GetBytes(Environment.GetEnvironmentVariable("JWT_KEY"));
        var tokenDescriptor = new SecurityTokenDescriptor
        {
            Subject = new ClaimsIdentity(new[]
            {
                new Claim(ClaimTypes.Name, user.Id.ToString()),
                new Claim(ClaimTypes.Email, user.Email)
            }),
            Expires = DateTime.UtcNow.AddHours(1),
            SigningCredentials = new SigningCredentials(new SymmetricSecurityKey(key), SecurityAlgorithms.HmacSha256Signature)
        };
        var token = tokenHandler.CreateToken(tokenDescriptor);
        return tokenHandler.WriteToken(token);
    }

    public bool VerifyPassword(string enteredPassword, string storedPasswordHash)
    {
        return BCrypt.Net.BCrypt.Verify(enteredPassword, storedPasswordHash);
    }
    
    private string HashPassword(string password)
    {
        return BCrypt.Net.BCrypt.HashPassword(password);
    }
}