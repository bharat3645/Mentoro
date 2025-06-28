#!/bin/bash
# Database Helper Scripts for Learning Buddy Platform

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Database connection settings
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-learning_buddy}
DB_USER=${DB_USER:-buddy_user}
DB_PASSWORD=${DB_PASSWORD:-buddy_pass}

# Function to check if PostgreSQL is running
check_postgres() {
    echo -e "${YELLOW}Checking PostgreSQL connection...${NC}"
    if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c '\q' 2>/dev/null; then
        echo -e "${GREEN}✓ PostgreSQL is running and accessible${NC}"
        return 0
    else
        echo -e "${RED}✗ Cannot connect to PostgreSQL${NC}"
        return 1
    fi
}

# Function to initialize database
init_db() {
    echo -e "${YELLOW}Initializing database...${NC}"
    if check_postgres; then
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f db/init.sql
        echo -e "${GREEN}✓ Database initialized successfully${NC}"
    else
        echo -e "${RED}✗ Failed to initialize database${NC}"
        exit 1
    fi
}

# Function to run migrations
run_migrations() {
    echo -e "${YELLOW}Running database migrations...${NC}"
    if check_postgres; then
        for migration in db/migrations/*.sql; do
            if [ -f "$migration" ]; then
                echo -e "${YELLOW}Running migration: $(basename $migration)${NC}"
                PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$migration"
            fi
        done
        echo -e "${GREEN}✓ All migrations completed${NC}"
    else
        echo -e "${RED}✗ Failed to run migrations${NC}"
        exit 1
    fi
}

# Function to seed test data
seed_data() {
    echo -e "${YELLOW}Seeding test data...${NC}"
    if check_postgres; then
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f db/seed.sql
        echo -e "${GREEN}✓ Test data seeded successfully${NC}"
    else
        echo -e "${RED}✗ Failed to seed test data${NC}"
        exit 1
    fi
}

# Function to reset database
reset_db() {
    echo -e "${YELLOW}Resetting database...${NC}"
    read -p "Are you sure you want to reset the database? This will delete all data. (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        if check_postgres; then
            PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
            init_db
            run_migrations
            seed_data
            echo -e "${GREEN}✓ Database reset completed${NC}"
        else
            echo -e "${RED}✗ Failed to reset database${NC}"
            exit 1
        fi
    else
        echo -e "${YELLOW}Database reset cancelled${NC}"
    fi
}

# Function to backup database
backup_db() {
    echo -e "${YELLOW}Creating database backup...${NC}"
    timestamp=$(date +"%Y%m%d_%H%M%S")
    backup_file="backup_${DB_NAME}_${timestamp}.sql"
    
    if check_postgres; then
        PGPASSWORD=$DB_PASSWORD pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME > "backups/$backup_file"
        echo -e "${GREEN}✓ Database backup created: backups/$backup_file${NC}"
    else
        echo -e "${RED}✗ Failed to create database backup${NC}"
        exit 1
    fi
}

# Function to show database status
show_status() {
    echo -e "${YELLOW}Database Status:${NC}"
    if check_postgres; then
        echo -e "${GREEN}Database: $DB_NAME${NC}"
        echo -e "${GREEN}Host: $DB_HOST:$DB_PORT${NC}"
        echo -e "${GREEN}User: $DB_USER${NC}"
        
        # Show table counts
        echo -e "\n${YELLOW}Table Statistics:${NC}"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
        SELECT 
            schemaname,
            tablename,
            n_tup_ins as inserts,
            n_tup_upd as updates,
            n_tup_del as deletes,
            n_live_tup as live_rows
        FROM pg_stat_user_tables 
        ORDER BY tablename;
        "
    fi
}

# Main script logic
case "$1" in
    "init")
        init_db
        ;;
    "migrate")
        run_migrations
        ;;
    "seed")
        seed_data
        ;;
    "reset")
        reset_db
        ;;
    "backup")
        mkdir -p backups
        backup_db
        ;;
    "status")
        show_status
        ;;
    "setup")
        init_db
        run_migrations
        seed_data
        ;;
    *)
        echo "Usage: $0 {init|migrate|seed|reset|backup|status|setup}"
        echo ""
        echo "Commands:"
        echo "  init     - Initialize database with schema"
        echo "  migrate  - Run database migrations"
        echo "  seed     - Seed database with test data"
        echo "  reset    - Reset database (WARNING: deletes all data)"
        echo "  backup   - Create database backup"
        echo "  status   - Show database status and statistics"
        echo "  setup    - Full setup (init + migrate + seed)"
        exit 1
        ;;
esac

