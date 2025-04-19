ALTER TABLE public.comments

ADD CONSTRAINT fk_comments_post
FOREIGN KEY (post_id) 
	REFERENCES public.posts(id);

ALTER TABLE public.comments
ADD CONSTRAINT fk_comments_user
FOREIGN KEY (user_id) 
	REFERENCES public.users(id);