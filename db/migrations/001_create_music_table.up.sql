CREATE TABLE public.music(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    song VARCHAR(100) NOT NULL,
    artist VARCHAR(100) NOT NULL,
    release_date DATE NOT NULL,
    link VARCHAR(100) NOT NULL,
    text TEXT NOT NULL
);