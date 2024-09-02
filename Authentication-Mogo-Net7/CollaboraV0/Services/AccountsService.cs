using CollaboraV0.DataTransferObjects;
using CollaboraV0.Interface;
using CollaboraV0.Models;
using DataTransferObjects;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.Extensions.Options;
using MongoDB.Driver;

namespace CollaboraV0.Services;

public class AccountsService : IAccountsServicecs
{
    private readonly IMongoCollection<UserModel> userCollection;
    private readonly MongoDbConnectionInfoHelper helperConnector;
    public AccountsService(IOptions<DBConnectionSettings>connection)
    {
        helperConnector = new MongoDbConnectionInfoHelper { CollectionName = "Users" , DatabaseName= "Accounts"};
        var mongoClient = new MongoClient(connection.Value.ConnectionString);
        var mongoDatabase = mongoClient.GetDatabase(helperConnector.DatabaseName);
        userCollection = mongoDatabase.GetCollection<UserModel>(helperConnector.CollectionName);
    }

    public async Task<UserModel?> createNewUser(RequestUserModel user)
    {
        var _user = new UserModel
        {
            EmailId = user.EmailId,
            Age = user.Age,
        };
        if (_user is null) return null;
        await userCollection.InsertOneAsync(_user);
        return _user;
    }

    public async  Task<UserModel?> login(string Email)
    {
        if (!string.IsNullOrEmpty(Email))
        {
            var res = await userCollection.Find(x => x.EmailId == Email).FirstOrDefaultAsync();
            return res;
        }
        return null;
    }
}
