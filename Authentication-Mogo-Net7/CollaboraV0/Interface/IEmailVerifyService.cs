namespace CollaboraV0.Interface
{
    public interface IEmailVerifyService
    {
        Task<bool?> verifyLogic(string Email);
    }
}
