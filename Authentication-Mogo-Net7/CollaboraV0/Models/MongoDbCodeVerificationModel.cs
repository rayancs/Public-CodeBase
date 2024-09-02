using DataTransferObjects;
using MongoDB.Bson.Serialization.Attributes;

namespace CollaboraV0.Models
{
    public class MongoDbCodeVerificationModel
    {

        [BsonId]
        [BsonRepresentation(MongoDB.Bson.BsonType.ObjectId)]
        public string? Id { get; set; }
        [BsonElement("Email")]
        public string? Email { get; set; } = null;
        [BsonElement("CreatedAt")]
        public DateTime? CreatedAt { get; set; }
        [BsonElement("ExpireAt")]
        public DateTime? ExpireAt { get; }
    }
}
