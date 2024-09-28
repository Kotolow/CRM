using System.Text.Json.Serialization;

namespace CRMAuth.Models;

public class ApiResponse<T>
{
    public int status_code { get; set; }
    public string message { get; set; }
    public T? data { get; set; }
}