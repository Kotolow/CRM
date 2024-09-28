using System.Security.Claims;
using CRMAuth.Models;
using CRMAuth.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace CRMAuth.Controllers;

[ApiController]
[Route("api/user")]
public class UserController : ControllerBase
{
    private readonly IUserService _userService;

    public UserController(IUserService userService)
    {
        _userService = userService;
    }
    
    [HttpGet("{id:int}")]
    [ProducesResponseType(typeof(ApiResponse<UserDto>), StatusCodes.Status200OK)]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status404NotFound)]
    [Produces("application/json")]
    public async Task<IActionResult> GetUserById([FromRoute] int id)
    {
        var response = await _userService.GetUserByIdAsync(id);
        if (response.status_code == 404)
            return NotFound(response);
        return Ok(response);
    }

    [HttpGet]
    [ProducesResponseType(typeof(ApiResponse<List<UserDto>>), StatusCodes.Status200OK)]
    [Produces("application/json")]
    private async Task<IActionResult> GetAllUsers()
    {
        
        var users = await _userService.GetAllUsersAsync();
        return Ok(users);
    }

    [HttpPost("login")]
    [Produces("application/json")]
    public async Task<IActionResult> Login([FromBody] LoginDto loginDto)
    {
        var user = await _userService.LoginAsync(loginDto);
        if (user == null)
        {
            return Unauthorized();
        }

        return Ok(new { Token = user.JwtToken });
    }
    
    [HttpPost("logout")]
    [Authorize]
    [Produces("application/json")]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status200OK)]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status401Unauthorized)]
    public async Task<IActionResult> Logout()
    {
        var userIdClaim = User.FindFirst(ClaimTypes.Name)?.Value;
        if (userIdClaim == null)
        {
            return Unauthorized(new ApiResponse<object>
            {
                status_code = 401,
                message = "Invalid token",
                data = null
            });
        }

        var userId = int.Parse(userIdClaim);
        await _userService.UpdateJwtTokenAsync(userId, null);
        return Ok(new ApiResponse<object>
        {
            status_code = 200,
            message = "Logout successful",
            data = null
        });
    }
    
    [HttpPost("register")]
    [ProducesResponseType(typeof(RegistrationDto), StatusCodes.Status201Created)]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status400BadRequest)]
    [Produces("application/json")]
    public async Task<IActionResult> CreateUser([FromBody] CreateUserDto user)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(new ApiResponse<object>
            {
                status_code = 400,
                message = "Invalid data",
                data = null
            });
        }

        var result = await _userService.CreateUserAsync(user);
        var registratedUser = await _userService.GetRegInfo(result.data);
        
        return CreatedAtAction(nameof(GetUserById), new { id = registratedUser.id }, registratedUser);
    }

    [HttpPost("upload-avatar")]
    [Produces("application/json")]
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status400BadRequest)]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status404NotFound)]
    public async Task<IActionResult> UploadAvatar([FromBody] UploadAvatarDto avatarDto)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(new ApiResponse<object>
            {
                status_code = 400,
                message = "Invalid input data",
                data = null
            });
        }
        
        var user = await _userService.GetUserInfoAsync(avatarDto.id);
        if (user == null)
        {
            return NotFound(new ApiResponse<object>
            {
                status_code = 404,
                message = "User not found",
                data = null
            });
        }

        var updatedUser = new UpdateUserDto
        {
            name = user.Name,
            avatar_url = avatarDto.avatar_url
        };
        
        await _userService.UpdateUserAsync(avatarDto.id, updatedUser);
        return Ok(new ApiResponse<object>
        {
            status_code = 201,
            message = "Avatar updated successfully",
            data = null
        });
    }
    
    [HttpPut("{id:int}")]
    [ProducesResponseType(typeof(ApiResponse<bool>), StatusCodes.Status200OK)]
    [ProducesResponseType(typeof(ApiResponse<object>), StatusCodes.Status404NotFound)]
    [Produces("application/json")]
    public async Task<IActionResult> UpdateUser(int id, [FromBody] UpdateUserDto user)
    {
        if (!ModelState.IsValid)
        {
            return BadRequest(new ApiResponse<object>
            {
                status_code = 400,
                message = "Invalid data",
                data = null
            });
        }

        var result = await _userService.UpdateUserAsync(id, user);
        if (result.status_code == 404)
        {
            return NotFound(result);
        }

        return Ok(result);
    }
}