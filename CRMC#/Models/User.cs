using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace CRMAuth.Models;

[Table("users")]
public class User
{
    [Key]
    [Column("id")]
    public int Id { get; set; }
    
    [Column("name")]
    [MaxLength(100)]
    public string Name { get; set; }
    
    [Column("email")]
    [MaxLength(100)]
    public string Email { get; set; }
    
    [Column("password_hash")]
    [MaxLength(450)]
    public string PasswordHash { get; set; }
    
    [Column("avatar_url")]
    [MaxLength(450)]
    public string? AvatarUrl { get; set; }
    
    [Column("google_calendar_token")]
    [MaxLength(450)]
    public string? GoogleCalendarToken { get; set; }
    
    [Column("created_at")]
    public DateTime CreatedAt { get; set; }
    
    [Column("updated_at")]
    public DateTime? UpdatedAt { get; set; }
    
    [Column("jwt_token")]
    [MaxLength(255)]
    public string? JwtToken { get; set; }
}