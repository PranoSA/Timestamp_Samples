

using Microsoft.EntityFrameworkCore;

namespace Database_Context {

    public class DB_Context : DbContext {
        public DbSet<Models.ModelWithTimestamp> ModelWithTimestamps { get; set; }
        public DbSet<Models.Notification> Notifications { get; set; }

        public DB_Context(DbContextOptions<DB_Context> options) : base(options) {
        }
        

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Models.ModelWithTimestamp>()
            .Property(e=>e.Id)
            .HasDefaultValueSql("gen_random_uuid()");

            modelBuilder.Entity<Models.Notification>()
            .Property(e=>e.Id)
            .HasDefaultValueSql("gen_random_uuid()");

        
        }


    }
}