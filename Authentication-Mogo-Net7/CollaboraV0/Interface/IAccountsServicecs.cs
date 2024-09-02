using CollaboraV0.DataTransferObjects;
using CollaboraV0.Models;
using DataTransferObjects;

namespace CollaboraV0.Interface;

public interface IAccountsServicecs
{
    Task<UserModel?> login(string Email);
    Task<UserModel?> createNewUser(RequestUserModel user);
}
