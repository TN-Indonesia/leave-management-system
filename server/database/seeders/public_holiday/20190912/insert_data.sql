INSERT INTO public_holiday
VALUES
  (1,'01-01-2019','01-01-2019','New Years Day'),
	(2,'05-02-2019','05-02-2019','Chinese New Year'),
	(3,'07-03-2019','07-03-2019','Bali Hindu New Year'),
	(4,'03-04-2019','03-04-2019','Isra Miraj'),
	(5,'19-04-2019','19-04-2019','Good Friday'),
	(6,'01-05-2019','01-05-2019','Labour Day'),
	(7,'19-05-2019','19-05-2019','Waisak Day'),
	(8,'30-05-2019','30-05-2019','Ascension Day of Jesus Christ'),
	(9,'01-06-2019','01-06-2019','Pancasila Day'),
	(10,'03-06-2019','04-06-2019','Lebaran Holiday'),
	(11,'05-06-2019','06-06-2019','Hari Raya Idul Fitri'),
	(12,'07-06-2019','07-06-2019','Lebaran Holiday'),
	(13,'11-08-2019','11-08-2019','Idul Adha'),
	(14,'17-08-2019','17-08-2019','Independence Day'),
	(15,'01-09-2019','01-09-2019','Islamic New Year'),
	(16,'09-11-2019','09-11-2019','Prophet Muhammads Birthday'),
	(17,'24-12-2019','24-12-2019','Christmas Holiday'),
	(18,'25-12-2019','25-12-2019','Christmas Day')
ON CONFLICT (id)
DO NOTHING;