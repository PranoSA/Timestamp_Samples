

namespace Models {

    public class ModelWithTimestamp {

        //ID auto-generated UUID in Postgres
        public Guid Id { get; set; }
        public string Name { get; set; }

        //Time Type in C#??
        public DateTimeOffset Timestamp { get; set; }
        

    }

    public class ModelWithTimestampRequest {
        public string Name { get; set; }

        public DateTimeOffset Timestamp { get; set; }
    }

    public class Notification {

        // ID auto-generated UUID in Postgres
        public Guid Id { get; set; }

        public DateTimeOffset Timestamp { get; set; }

        public Int32 UnixTimestamp32 { get; set; }

        public Int64 UnixTimestamp64 { get; set; }
    }
}