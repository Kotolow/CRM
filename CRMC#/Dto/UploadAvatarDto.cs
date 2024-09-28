using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace CRMAuth.Models;

public class UploadAvatarDto
{
    [Required]
    public int id { get; set; }
    [Required]
    public string avatar_url { get; set; }
}