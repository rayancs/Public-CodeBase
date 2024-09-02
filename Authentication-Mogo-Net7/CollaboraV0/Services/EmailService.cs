using CollaboraV0.Interface;
using System.Net.Mail;
using System.Net;
using CollaboraV0.Models;
using Microsoft.Extensions.Options;

namespace CollaboraV0.Services
{
    public class EmailService : IEmailService
    {
        private readonly EmailPrimitiveNetworkCredentials _credentials;
        public EmailService(IOptions<EmailPrimitiveNetworkCredentials> credentials)
        {
   
            _credentials!.SmptClient = credentials.Value.SmptClient;
            _credentials!.Port = credentials.Value.Port;
            _credentials!.SenderEmail = credentials.Value.SenderEmail;
            _credentials!.SenderEmailPassworsd = credentials .Value.SenderEmailPassworsd;
        }
        public async Task<bool> SendEmailAsync(string Email,string subject ,string Cont)
        {
            SmtpClient client = new SmtpClient(_credentials.SmptClient);
            client.Port = 587; 
            client.Credentials = new NetworkCredential(_credentials.SenderEmail,_credentials.SenderEmailPassworsd); // Replace with your SMTP credentials
            client.EnableSsl = true;
            try
            {
                var mailMessage = new MailMessage{
                    Subject = subject, Body = Cont
                };
                await client.SendMailAsync(mailMessage);
                return true;

            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error sending email: {ex.Message}");
                return false;
            }
        }

        
    }
}
