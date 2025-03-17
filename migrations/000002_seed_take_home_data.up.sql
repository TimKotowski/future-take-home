-- Take home requried test data
insert into users (id, first_name, email)
values
  (1, 'test_user_one', 'test_user_one@gmail.com'),
  (2, 'test_user_two', 'test_user_two@gmail.com'),
  (3, 'test_user_three', 'test_user_three@gmail.com'),
  (4, 'test_user_four', 'test_user_four@gmail.com'),
  (5, 'test_user_five', 'test_user_five@gmail.com'),
  (6, 'test_user_six', 'test_user_six@gmail.com'),
  (7, 'test_user_seven', 'test_user_seven@gmail.com'),
  (8, 'test_user_eight', 'test_user_eight@gmail.com'),
  (9, 'test_user_nine', 'test_user_nine@gmail.com'),
  (10, 'test_user_ten', 'test_user_ten@gmail.com');

insert into trainers (id, first_name)
values
 (1, 'test_trainers_one'),
 (2, 'test_trainers_two'),
 (3, 'test_trainers_three');

insert into appointments (id, trainer_id, user_id, start_slot, end_slot, status)
values
 (gen_random_uuid(), 1, 1, '2019-01-24T09:00:00-08:00', '2019-01-24T09:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 1, 2, '2019-01-24T10:00:00-08:00', '2019-01-24T10:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 1, 3, '2019-01-25T10:00:00-08:00', '2019-01-25T10:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 1, 4, '2019-01-25T10:30:00-08:00', '2019-01-25T11:00:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 1, 5, '2019-01-26T10:00:00-08:00', '2019-01-26T10:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 2, 6, '2019-01-24T09:00:00-08:00', '2019-01-24T09:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 2, 7, '2019-01-26T10:00:00-08:00', '2019-01-26T10:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 3, 8, '2019-01-26T12:00:00-08:00', '2019-01-26T12:30:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 3, 9, '2019-01-26T13:00:00-08:00', '2019-01-26T14:00:00-08:00', 'ACTIVE'),
 (gen_random_uuid(), 3, 10, '2019-01-26T14:00:00-08:00', '2019-01-26T14:30:00-08:00', 'ACTIVE');
