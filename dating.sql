

CREATE TABLE public.profile (
    id uuid NOT NULL,
    name character varying,
    age bigint,
    sex character varying,
    orientation character varying,
    status character varying,
    account character varying,
    email character varying,
    password character varying,
    created_at date
);


--
-- TOC entry 203 (class 1259 OID 81137)
-- Name: profile_preferences; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profile_preferences (
    id uuid NOT NULL,
    people_id uuid,
    "like" uuid[],
    pass uuid[],
    view bigint,
    jwt character varying
);


--
-- TOC entry 204 (class 1259 OID 89321)
-- Name: verification; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.verification (
    id uuid NOT NULL,
    people_id uuid,
    expiresat date,
    verified boolean DEFAULT false
);


--
-- TOC entry 2830 (class 0 OID 81129)
-- Dependencies: 202
-- Data for Name: profile; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.profile VALUES ('6b7ebb4b-c0eb-4fa3-a8d2-a8cd36d0f47f', 'melina', 22, 'female', 'straight', 'single', 'free', 'melina@melina.com', '$2a$14$L9rvisRkcLZ1XrqFkNs9i.kJ1aucb9xmHgMFpHrG7pz.2ygVU1m36', '2023-07-31');
INSERT INTO public.profile VALUES ('edd196cd-919d-4292-94ce-5044b04c1942', 'dikan', 22, 'male', 'straight', 'single', 'free', 'dikan@dikan.com', '$2a$14$5PI69r5aJx4/pM87KhUnsO2Swkhpaf6PAorIZ0cOHz6IdEyGkTPW2', '2023-07-31');
INSERT INTO public.profile VALUES ('d501b573-b01c-4c5c-8339-46ab37549542', 'dikano', 22, 'male', 'straight', 'single', 'free', 'dikano@dikan.com', '$2a$14$nhqgwuObrld/LTNaqUESjurW8kDh1J893Wv0kxN2wDpXgRxuh5BzW', '2023-07-31');
INSERT INTO public.profile VALUES ('e7c1b691-59e5-4e9a-8b3b-b63d628ca984', 'dikanor', 22, 'male', 'straight', 'single', 'free', 'dikanor@dikan.com', '$2a$14$iVj49LHM6Lg3n.uglPCAB.JvKsEp1v4pHFLq1B8HZ9G7oGg2pQE7y', '2023-07-31');
INSERT INTO public.profile VALUES ('4f3b7289-f727-4fe2-a7c7-25907e0e9553', 'melinan', 22, 'female', 'straight', 'single', 'free', 'melinan@melinan.com', '$2a$14$Ya0pMzBxIH3buBKVeBqrHOFKaI4R3Xw8/7s.W6yQF.cjJFHz5ZIRW', '2023-07-31');
INSERT INTO public.profile VALUES ('6274f776-610e-4867-b472-855e287d5abd', 'melinano', 22, 'female', 'straight', 'single', 'free', 'melinano@melinano.com', '$2a$14$0GqpgYSubR3hwl0Dtu7REebQbFuwGzfcYostX/WQExKKckTTVBMk6', '2023-07-31');
INSERT INTO public.profile VALUES ('5cd67570-581e-44b3-a1d2-fda9be9324c3', 'melinanoz', 22, 'female', 'straight', 'single', 'free', 'melinanoz@melinanoz.com', '$2a$14$xPIMwpI24D9elWVcRkw4.OcjsryFro69bIjJZ8TA.vlWG1g/t1WlW', '2023-07-31');
INSERT INTO public.profile VALUES ('b829440e-5b0f-4539-af02-e6edc5b332ad', 'melinanozi', 22, 'female', 'straight', 'single', 'free', 'melinanozi@melinanozi.com', '$2a$14$sEymTzsuqCtjsyyL3Et6Ye5QI4nCn6J7rlozH//fujbtIDMO9.Wou', '2023-07-31');
INSERT INTO public.profile VALUES ('6d0d7aaf-75f1-49fe-8f1f-94096bfad0e0', 'melinanozik', 22, 'female', 'straight', 'single', 'free', 'melinanozik@melinanozik.com', '$2a$14$8sAnaSTibPFpNuZugTNxmeXUHp6Xtor7V8GYE4qnHlMilDQDBZmAy', '2023-07-31');
INSERT INTO public.profile VALUES ('4b2d6a52-b767-41c6-b89b-841dab3cd70a', 'melinanoziki', 22, 'female', 'straight', 'single', 'free', 'melinanoziki@melinanoziki.com', '$2a$14$FMBksdYT/OZR5Pr44iKhC.icjgRAdOymj5SimkUY3QbBvVhK2W0am', '2023-07-31');
INSERT INTO public.profile VALUES ('a8e8ea61-c970-41a5-911b-11dc3e5d2770', 'melinanozikiy', 22, 'female', 'straight', 'single', 'free', 'melinanozikiy@melinanozikiy.com', '$2a$14$WHCgEUMP1PRUsxupwn7TcOulY4.J1a1o1Vd0E/D/DeYBKGvKyZoja', '2023-07-31');
INSERT INTO public.profile VALUES ('0f8072e4-39d9-42f5-ac55-ae951275040c', 'melinanozikiyk', 22, 'female', 'straight', 'single', 'free', 'melinanozikiyk@melinanozikiyk.com', '$2a$14$rpesRC3IfTvEYSvcf78/Oeu0nRABJVwsp6D1s7pwQzNuyLEwzeb26', '2023-07-31');
INSERT INTO public.profile VALUES ('25891148-8413-4603-900d-91a5deeb47fc', 'melinanozikiykl', 22, 'female', 'straight', 'single', 'free', 'melinanozikiykl@melinanozikiykl.com', '$2a$14$XNrBGYyBits1l4LxrF.FrOEuZWYIKa8HRjrGmb/40CHA9JevicSKi', '2023-07-31');
INSERT INTO public.profile VALUES ('7e24ef9c-0bb8-43f2-b416-7a1a9b805880', 'dikanori', 22, 'male', 'straight', 'single', 'free', 'dikanori@dikanori.com', '$2a$14$A3rTj266BDnldtGCBMljk.TjUiWwjctHkOYocTdm1omHV.mqoIGpa', '2023-07-31');
INSERT INTO public.profile VALUES ('209b0712-944a-46cd-972d-f92122a7fd88', 'dikanorim', 22, 'male', 'straight', 'single', 'free', 'dikanorim@dikanorim.com', '$2a$14$vq4AgWLEsFFpHIZQ3eVB6u0FeOrDNcdb/XqOoyuT6oG7DUPeBA8Zu', '2023-07-31');
INSERT INTO public.profile VALUES ('786cdf1c-1b73-4764-991d-4dbc046d6ffc', 'dikanorimo', 22, 'male', 'straight', 'single', 'free', 'dikanorimo@dikanorimo.com', '$2a$14$n8BIV8yeqsNYNtHq.ISPe./siVrBI4sTNP6GC313lMiRT9Zj2B1VC', '2023-07-31');
INSERT INTO public.profile VALUES ('ec82e5c5-6244-4da0-ac0c-971f3138545a', 'dikanorimop', 22, 'male', 'straight', 'single', 'free', 'dikanorimop@dikanorimop.com', '$2a$14$ku6z6d.zTYMYORV7U4jcSes6IwyfFc9xMqtkxOrAo4xpr6E7z4Uta', '2023-07-31');
INSERT INTO public.profile VALUES ('92cd6939-476e-46ac-9044-71dacaff27d7', 'dikanorimopl', 22, 'male', 'straight', 'single', 'free', 'dikanorimopl@dikanorimopl.com', '$2a$14$GKgV3jhpLBI/3wwaJYRAXuaHuZb8g3Qq0Ap.4P4EQsO01QlUqn.q.', '2023-07-31');
INSERT INTO public.profile VALUES ('0a1248fb-6c35-4844-a1d5-2c3d9887c638', 'dikanorimoplp', 22, 'male', 'straight', 'single', 'free', 'dikanorimoplp@dikanorimoplp.com', '$2a$14$wIX7qJYXXg5P/YHKy1P/4eafmDEYO1ERBg1SL5EAJQCtw05GOUvB6', '2023-07-31');
INSERT INTO public.profile VALUES ('772d7926-8065-43ce-9c16-272c5ffd50ee', 'dika', 22, 'male', 'straight', 'single', 'free', 'dika@dika.com', '$2a$14$6SKoeL718gG2fMnAxegqF.EblhOAFBLGykp5k2GnQkExe9VlGkih6', '2023-07-31');


--
-- TOC entry 2831 (class 0 OID 81137)
-- Dependencies: 203
-- Data for Name: profile_preferences; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.profile_preferences VALUES ('151a56f6-d914-4f4e-bfe7-41d3457282a5', '6b7ebb4b-c0eb-4fa3-a8d2-a8cd36d0f47f', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('56e35748-f7ca-4c22-8126-532326e9d465', 'edd196cd-919d-4292-94ce-5044b04c1942', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('6f7ba862-b11e-4369-8fed-3019b71048dc', 'd501b573-b01c-4c5c-8339-46ab37549542', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('d7d1c1ec-08e7-4bd7-8bc5-84d8bd25c9c0', 'e7c1b691-59e5-4e9a-8b3b-b63d628ca984', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('c3869293-8e8e-45fd-92e7-91847cce4fa8', '4f3b7289-f727-4fe2-a7c7-25907e0e9553', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('4a35a6c2-1d23-4464-bada-035b1a11aebb', '6274f776-610e-4867-b472-855e287d5abd', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('4e1c076f-4123-425e-ab1d-adaf5b19d047', '5cd67570-581e-44b3-a1d2-fda9be9324c3', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('22e7dab8-d3d3-4343-be70-5ddbbfba97fa', 'b829440e-5b0f-4539-af02-e6edc5b332ad', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('b65b6244-d3d4-48d9-b105-04df43cd2cad', '6d0d7aaf-75f1-49fe-8f1f-94096bfad0e0', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('56c5f870-4034-4898-a11c-46c05f4da1af', '4b2d6a52-b767-41c6-b89b-841dab3cd70a', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('df21b56e-9772-49c0-9b06-f451e240fb0a', 'a8e8ea61-c970-41a5-911b-11dc3e5d2770', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('f798c8b8-766a-4173-8bff-2c3b4975c25b', '0f8072e4-39d9-42f5-ac55-ae951275040c', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('f3e7ebbd-ba46-422c-a736-e41bca36cd93', '25891148-8413-4603-900d-91a5deeb47fc', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('867be25e-8647-41f7-9257-69eccea8fdf9', '7e24ef9c-0bb8-43f2-b416-7a1a9b805880', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('17beb7ec-341a-4db6-86b6-ead6e707e28d', '209b0712-944a-46cd-972d-f92122a7fd88', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('778af1da-97b4-4637-95ba-7d38778940f6', '786cdf1c-1b73-4764-991d-4dbc046d6ffc', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('3de43b00-b5c9-4b79-bc8e-2ae8d7e33043', 'ec82e5c5-6244-4da0-ac0c-971f3138545a', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('0623eb38-02bd-43ad-9748-090f39f2e64f', '92cd6939-476e-46ac-9044-71dacaff27d7', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('abda37ae-8b69-4c2e-afa5-be6188e5b2a2', '0a1248fb-6c35-4844-a1d5-2c3d9887c638', NULL, NULL, 0, NULL);
INSERT INTO public.profile_preferences VALUES ('8a12bfa4-3401-4289-8cbc-6a614ef249c9', '772d7926-8065-43ce-9c16-272c5ffd50ee', '{4b2d6a52-b767-41c6-b89b-841dab3cd70a,25891148-8413-4603-900d-91a5deeb47fc}', '{b829440e-5b0f-4539-af02-e6edc5b332ad,4f3b7289-f727-4fe2-a7c7-25907e0e9553,25891148-8413-4603-900d-91a5deeb47fc}', 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc3MmQ3OTI2LTgwNjUtNDNjZS05YzE2LTI3MmM1ZmZkNTBlZSIsIm5hbWUiOiJkaWthIiwiYWdlIjoyMiwic2V4IjoibWFsZSIsIm9yaWVudGF0aW9uIjoic3RyYWlnaHQiLCJzdGF0dXMiOiJzaW5nbGUiLCJhY2NvdW50IjoicHJlbWl1bSIsImVtYWlsIjoiZGlrYUBkaWthLmNvbSIsInBhc3N3b3JkIjoiJDJhJDE0JDZTS29lTDcxOGdHMmZNbkF4ZWdxRi5FYmxoT0FGQkxHeWtwNWsyR25Ra0V4ZTlWbEdraWg2IiwibGFzdF9sb2dpbiI6IjIwMjMtMDgtMDFUMjE6Mjg6MzguMzAyNjI1OCswNzowMCIsInZlcmlmaWVkIjp0cnVlLCJleHAiOjE2OTA5ODY1MTh9.FSsV-9kCIbApz0FuIF9ftDMuyVLi_BGDi0uKOe0yCIo');


--
-- TOC entry 2832 (class 0 OID 89321)
-- Dependencies: 204
-- Data for Name: verification; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.verification VALUES ('d51612ed-752e-4fc9-8efe-1e78a65ea3d1', '6b7ebb4b-c0eb-4fa3-a8d2-a8cd36d0f47f', '2023-08-01', true);
INSERT INTO public.verification VALUES ('8d569bbe-2ee0-4aa4-b671-fb75d7a1c204', '772d7926-8065-43ce-9c16-272c5ffd50ee', '2023-08-01', true);
INSERT INTO public.verification VALUES ('a7a908ec-2a3f-40fa-8f3f-c8b902320f89', 'edd196cd-919d-4292-94ce-5044b04c1942', '2023-08-01', true);
INSERT INTO public.verification VALUES ('e682d9ce-0326-4004-aea7-fcb03634cde6', 'd501b573-b01c-4c5c-8339-46ab37549542', '2023-08-01', true);
INSERT INTO public.verification VALUES ('ffe4a0a1-bcfb-400a-a951-0ef8ffba4625', 'e7c1b691-59e5-4e9a-8b3b-b63d628ca984', '2023-08-01', true);
INSERT INTO public.verification VALUES ('c0732e06-8bf9-41ea-b67b-819009afd9ee', '4f3b7289-f727-4fe2-a7c7-25907e0e9553', '2023-08-01', true);
INSERT INTO public.verification VALUES ('22eadb3e-ed21-48db-b293-33095e926c9a', '6274f776-610e-4867-b472-855e287d5abd', '2023-08-01', true);
INSERT INTO public.verification VALUES ('b3b9cbe8-e563-4154-9463-0a114019f148', '5cd67570-581e-44b3-a1d2-fda9be9324c3', '2023-08-01', true);
INSERT INTO public.verification VALUES ('fe3557d7-717c-4eb0-b51e-c16390ac7a90', 'b829440e-5b0f-4539-af02-e6edc5b332ad', '2023-08-01', true);
INSERT INTO public.verification VALUES ('b3123780-3ef6-48c9-b96a-6bcec57609f9', '6d0d7aaf-75f1-49fe-8f1f-94096bfad0e0', '2023-08-01', true);
INSERT INTO public.verification VALUES ('60f3d06b-09c3-4503-bc06-47c32827c08e', '4b2d6a52-b767-41c6-b89b-841dab3cd70a', '2023-08-01', true);
INSERT INTO public.verification VALUES ('2f1b1dd5-3af2-45e9-8b04-6bd9d2e5762c', 'a8e8ea61-c970-41a5-911b-11dc3e5d2770', '2023-08-01', true);
INSERT INTO public.verification VALUES ('8574d5d0-62c1-4048-bb8a-8fa6a5e5264b', '0f8072e4-39d9-42f5-ac55-ae951275040c', '2023-08-01', true);
INSERT INTO public.verification VALUES ('8053ec58-a5b4-4998-a4e3-6223040f8b8a', '25891148-8413-4603-900d-91a5deeb47fc', '2023-08-01', true);
INSERT INTO public.verification VALUES ('0fa72534-f724-40e9-95d3-3c3b698f32d3', '7e24ef9c-0bb8-43f2-b416-7a1a9b805880', '2023-08-01', true);
INSERT INTO public.verification VALUES ('f9dace66-5339-451a-99eb-d238d5a9b305', '209b0712-944a-46cd-972d-f92122a7fd88', '2023-08-01', true);
INSERT INTO public.verification VALUES ('21191858-3757-4f8c-aa9e-ee5374f94f91', '786cdf1c-1b73-4764-991d-4dbc046d6ffc', '2023-08-01', true);
INSERT INTO public.verification VALUES ('173212e8-9c49-4cb9-bb32-602383edea15', 'ec82e5c5-6244-4da0-ac0c-971f3138545a', '2023-08-01', true);
INSERT INTO public.verification VALUES ('b4c4eba7-eaf0-4fa8-8007-11b9f774380f', '92cd6939-476e-46ac-9044-71dacaff27d7', '2023-08-01', true);
INSERT INTO public.verification VALUES ('94200df5-6e83-4dbe-8795-e6f31e743a8f', '0a1248fb-6c35-4844-a1d5-2c3d9887c638', '2023-08-01', true);


--
-- TOC entry 2697 (class 2606 OID 81136)
-- Name: profile profile_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile
    ADD CONSTRAINT profile_pkey PRIMARY KEY (id);


--
-- TOC entry 2699 (class 2606 OID 81144)
-- Name: profile_preferences profile_preferences_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_preferences
    ADD CONSTRAINT profile_preferences_pkey PRIMARY KEY (id);


--
-- TOC entry 2701 (class 2606 OID 89326)
-- Name: verification verification_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verification
    ADD CONSTRAINT verification_pkey PRIMARY KEY (id);


--
-- TOC entry 2702 (class 2606 OID 81145)
-- Name: profile_preferences fk_profile_preferences_profile; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_preferences
    ADD CONSTRAINT fk_profile_preferences_profile FOREIGN KEY (people_id) REFERENCES public.profile(id);


--
-- TOC entry 2703 (class 2606 OID 89327)
-- Name: verification fk_verification_profile; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verification
    ADD CONSTRAINT fk_verification_profile FOREIGN KEY (people_id) REFERENCES public.profile(id);


-- Completed on 2023-08-01 21:39:21

--
-- PostgreSQL database dump complete
--

