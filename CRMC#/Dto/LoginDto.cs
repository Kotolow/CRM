using System.ComponentModel.DataAnnotations;
using Newtonsoft.Json;

namespace CRMAuth.Models;

public class LoginDto
{
    [Required]
    [EmailAddress]
    [JsonProperty("email")]
    public string Email { get; set; }
    
    [Required]
    [JsonProperty("password")]
    public string Password { get; set; }
}

public class LoginResponseDto
{
    [Required]
    [JsonProperty("jwt_token")]
    public string JwtToken { get; set; }
}