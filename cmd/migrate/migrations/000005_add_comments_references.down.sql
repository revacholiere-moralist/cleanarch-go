ALTER TABLE public.comments
REMOVE CONSTRAINT fk_comments_post;

ALTER TABLE public.comments
REMOVE CONSTRAINT fk_comments_user
