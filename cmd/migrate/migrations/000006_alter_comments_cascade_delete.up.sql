ALTER TABLE public.comments

DROP CONSTRAINT fk_comments_post,
ADD CONSTRAINT fk_comments_post
FOREIGN KEY (post_id) 
	REFERENCES public.posts(id) 
    ON DELETE CASCADE;

ALTER TABLE public.comments
DROP CONSTRAINT fk_comments_user,
ADD CONSTRAINT fk_comments_user
FOREIGN KEY (user_id) 
	REFERENCES public.users(id)
    ON DELETE CASCADE;
