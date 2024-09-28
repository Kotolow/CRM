using System.ComponentModel.DataAnnotations;
using Newtonsoft.Json;

namespace CRMAuth.Models;

public class RegistrationDto
{
    [Required]
    public int id { get; set; }
    [Required]
    public string name { get; set; }
    [Required]
    [EmailAddress]
    public string email { get; set; }
    public string? avatar_url { get; set; }
    [Required]
    public DateTime created_at { get; set; } 
}