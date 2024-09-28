using CRMAuth.Models;
using Microsoft.EntityFrameworkCore;

namespace CRMAuth.Data;

public class UserRepository : IUserRepository
{
    private readonly ApplicationDbContext _context;
    
    public UserRepository(ApplicationDbContext context)
    {
        _context = context;
    }

    public async Task<User> GetUserByIdAsync(int id)
    {
        return await _context.Users.FindAsync(id);
    }

    public async Task AddUserAsync(User user)
    {
        await _context.Users.AddAsync(user);
        await _context.SaveChangesAsync();
    }

    public async Task UpdateUserAsync(User user)
    {
        var existingUser = await _context.Users.FindAsync(user.Id);
        if (existingUser == null) return; 

        existingUser.Name = user.Name;
        existingUser.AvatarUrl = user.AvatarUrl;
        
        existingUser.UpdatedAt = DateTime.UtcNow;
        existingUser.CreatedAt = existingUser.CreatedAt.Kind == DateTimeKind.Unspecified
            ? DateTime.SpecifyKind(existingUser.CreatedAt, DateTimeKind.Utc)
            : existingUser.CreatedAt.ToUniversalTime();
        
        await _context.SaveChangesAsync();
    }

    public async Task UpdateJwtTokenAsync(int userId, string token)
    {
        var user = await _context.Users.FindAsync(userId);
        if (user != null)
        {
            user.JwtToken = token;
            await _context.SaveChangesAsync();
        }
    }

    public async Task<List<User>> GetAllUsersAsync()
    {
        return await _context.Users.ToListAsync();
    }
}