namespace CollaboraV0.Interface;

public interface IEmailService
{
    Task<bool> SendEmailAsync(string Email ,string subject , string Cont);
}
