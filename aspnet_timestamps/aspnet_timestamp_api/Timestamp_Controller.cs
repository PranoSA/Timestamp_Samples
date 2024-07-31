


using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Models;

namespace Controllers {

    // REST CONTROLLER Timestamp
    [Route("api/[controller]")]
    [ApiController]
    public class TimestampController: ControllerBase {

        private Database_Context.DB_Context _context;

        public TimestampController(Database_Context.DB_Context context) {
            _context = context;
        }

        // GET api/timestamp
        [HttpGet]
        public ActionResult<List<ModelWithTimestamp>> Get() {
            // Get DB Context
            
            // Get all timetstamps
            var timestamps = _context.ModelWithTimestamps.ToList();

            // Return all timestamp 
            return timestamps;

        }

        // POST api/timestamp
        [HttpPost]
        public async Task<ActionResult<string>> Post([FromBody] Models.ModelWithTimestampRequest model) {
            
            // Create new ModelWithTimestamp
            var newModel = new Models.ModelWithTimestamp {
                Name = model.Name,
                Timestamp = model.Timestamp
            };

            newModel.Timestamp = newModel.Timestamp.ToUniversalTime(); // Convert to UTC

            // Add to DB Context
            _context.ModelWithTimestamps.Add(newModel);

            // Save Changes
           await  _context.SaveChangesAsync();

            // Subtract 1 hour, 30 minutes, 2 seconds, and 500 milliseconds from the timestamp
           var TimeForNotificaition = newModel.Timestamp.Subtract(new TimeSpan(0,1,30,2,500));



            var timeInt32 = (Int32) TimeForNotificaition.ToUnixTimeSeconds();
            var timeInt64 = TimeForNotificaition.ToUnixTimeMilliseconds();

           //Create New Notification 
              var newNotification = new Models.Notification {
                Timestamp = newModel.Timestamp,
                UnixTimestamp32 = timeInt32,
                UnixTimestamp64 = timeInt64
            };

            /// Add to DB Context
            _context.Notifications.Add(newNotification);

            // Save Changes
            await _context.SaveChangesAsync();

            //Create Notification 

            return "Timestamp Controller";
        }

        // PUT api/timestamp
        [HttpPut]
        public ActionResult<string> Put([FromBody] Models.ModelWithTimestampRequest model) {
            return "Timestamp Controller";
        }

        // DELETE api/timestamp
        [HttpDelete]
        public ActionResult<string> Delete() {
            return "Timestamp Controller";
        }
    }

    //Notification Controllers

    [Route("api/[controller]")]
    [ApiController]
    public class NotificationController: ControllerBase {

        private Database_Context.DB_Context _context;

        public NotificationController(Database_Context.DB_Context context) {
            _context = context;
        }

        // GET api/notification
        [HttpGet]
        public async Task<ActionResult<List<Notification>>> Get() {
            // Get DB Context
            
            // Get all timetstamps
            var notifications = await _context.Notifications.ToListAsync();

            // Return all timestamp 
            return notifications;

        }

        // POST api/notification
        [HttpPost]
        public ActionResult<string> Post([FromBody] Models.Notification model) {
            return "Notification Controller";
        }

        // PUT api/notification
        [HttpPut]
        public ActionResult<string> Put([FromBody] Models.Notification model) {
            return "Notification Controller";
        }

        // DELETE api/notification
        [HttpDelete]
        public ActionResult<string> Delete() {
            return "Notification Controller";
        }
    }

}