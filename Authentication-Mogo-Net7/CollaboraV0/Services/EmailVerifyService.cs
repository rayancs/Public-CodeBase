using CollaboraV0.Interface;

namespace CollaboraV0.Services;

public class EmailVerifyService : IEmailVerifyService
{
    private readonly IEmailService emailService;
    public EmailVerifyService(IEmailService emailService)
    {
        this.emailService = emailService;
    }

    public async Task<bool?> verifyLogic(string Email)
    {
        var subject = @"Verification Code,noreply";
        var body = "<h1>Do Not share</h1> <p>code:</p>";
        var res = await emailService.SendEmailAsync(Email,subject,body)
    }
}
