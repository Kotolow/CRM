using System.ComponentModel.DataAnnotations;
using Newtonsoft.Json;

namespace CRMAuth.Models;

public class UserDto
{
    public string name { get; set; }
    public string email { get; set; }
    public string? avatar_url { get; set; }
}

public class CreateUserDto
{
    [Required]
    [MaxLength(100)]
    public string name { get; set; }
    [Required]
    [EmailAddress]
    [MaxLength(100)]
    public string email { get; set; }
    [Required]
    [MaxLength(450)]
    public string password { get; set; }
    [MaxLength(450)]
    public string? avatar_url { get; set; }
}

public class UpdateUserDto
{
    [MaxLength(100)]
    public string name { get; set; }
    [MaxLength(450)]
    public string? avatar_url { get; set; }
    // [EmailAddress]
    // [MaxLength(100)]
    // public string Email { get; set; }
    // [MaxLength(450)]
    // public string Password { get; set; } // to add next time
}