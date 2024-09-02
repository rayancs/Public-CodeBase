using System.Net.Mail;

namespace CollaboraV0.Models;

public class EmailPrimitiveNetworkCredentials
{
    public string?SmptClient { get; set; } 
    public int? Port { get; set; } = 587;
    public string? SenderEmail { get; set; }
    public string? SenderEmailPassworsd { get ; set; }
    

}
