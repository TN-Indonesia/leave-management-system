CREATE TABLE IF NOT EXISTS public_holiday 
    ( 
    "id" SERIAL PRIMARY KEY, 
    "date_start" TEXT NOT NULL, 
    "date_end" TEXT NOT NULL, 
    "description" TEXT NOT NULL 
    );