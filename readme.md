#This is the clean code architecture of go used to write microservices
#How to build and run

1. Clone the repo 
2. Change config.yml file to change you config details
3. Run "make proto" command. It will all binary files in "usecase" folder
4. Run "make grpc" command. 
5. Run "make rest" command. 
5. Run  "./bin/grpc --server_address=127.0.0.1:9090"
7. Open new  terminal and run "./bin/rest". You can see server is running on port 8080
8. Hit the rest API on localhost:8080/user/getall

##For more
1. Create statement to create users table
	CREATE TABLE `users` (
	  `id` int(11) NOT NULL,
	  `email_id` varchar(45) DEFAULT NULL,
	  `password` varchar(45) DEFAULT NULL,
	  `created_at` timestamp(6) NULL DEFAULT NULL,
	  `updated_at` timestamp(6) NULL DEFAULT NULL,
	  `token` varchar(200) DEFAULT NULL,
	  `name` varchar(255) DEFAULT NULL,
	  `xxx_unrecognized` varbinary(255) DEFAULT NULL,
	  `xxx_sizecache` int(11) DEFAULT NULL,
	  PRIMARY KEY (`id`)
	) ENGINE=InnoDB;

