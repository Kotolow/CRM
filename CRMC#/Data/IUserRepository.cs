using CRMAuth.Models;

namespace CRMAuth.Data;

public interface IUserRepository
{
    Task<User> GetUserByIdAsync(int id);
    Task AddUserAsync(User user);
    Task UpdateUserAsync(User user);
    Task<List<User>> GetAllUsersAsync();
    Task UpdateJwtTokenAsync(int userId, string token);
}