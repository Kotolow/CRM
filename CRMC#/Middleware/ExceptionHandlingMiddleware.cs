using CRMAuth.Models;
using Newtonsoft.Json;

namespace CRMAuth.Middleware;

public class ExceptionHandlingMiddleware
{
    private readonly RequestDelegate _requestDelegate;

    public ExceptionHandlingMiddleware(RequestDelegate requestDelegate)
    {
        _requestDelegate = requestDelegate;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        try
        {
            await _requestDelegate(context);
        }
        catch (Exception e)
        {
            await HandleExceptionAsync(context, e);
        }
    }

    private Task HandleExceptionAsync(HttpContext context, Exception ex)
    {
        context.Response.ContentType = "application/json";
        context.Response.StatusCode = StatusCodes.Status500InternalServerError;
        var response = new ApiResponse<object>
        {
            status_code = context.Response.StatusCode,
            message = "Server internal error. Please try later!",
            data = null
        };

        return context.Response.WriteAsync(JsonConvert.SerializeObject(response));
    }
}