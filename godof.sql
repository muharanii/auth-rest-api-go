PGDMP  0    2                |            godof    16.3    16.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            	           1262    16398    godof    DATABASE     g   CREATE DATABASE godof WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE godof;
                postgres    false            �            1259    16414    users    TABLE     �   CREATE TABLE public.users (
    id integer NOT NULL,
    full_name character varying(255) NOT NULL,
    phone character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false                      0    16414    users 
   TABLE DATA           I   COPY public.users (id, full_name, phone, username, password) FROM stdin;
    public          postgres    false    215   �       s           2606    16420    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    215               ]   x�3��-�H,J���4�0� N_���Q�����~P�W��kziQZv�Q���YrPjjrx��E�i��^v��Ovf��{R`�A���AA��W� M �     