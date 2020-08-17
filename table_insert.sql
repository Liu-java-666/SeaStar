/*
SQLyog 企业版 - MySQL GUI v8.14
MySQL - 5.7.26 : Database - vparty
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

USE `vparty`;

/*Data for the table `action_log` */

/*Data for the table `blacklist` */

/*Data for the table `captcha` */

/*Data for the table `certification` */

/*Data for the table `denounce` */

/*Data for the table `dynamic` */

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (1,11,'2020-04-02 18:20:00','带爸妈自驾莫干山','','image','46,47,48,49,50,51',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (2,12,'2020-04-20 20:00:00','你很可爱 不是cute，是could be loved','','image','52,53,54',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (3,13,'2020-03-28 14:42:00','happy birthday','','image','55,56,57',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (4,14,'2020-01-28 22:42:00','职场穿搭','','image','58,59,60,61,62,63',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (5,15,'2020-03-31 17:15:00','不断进步，才不至于退步！','','image','64,65,66',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (6,16,'2020-03-30 17:00:00','老片计划#2017#','','image','67,68,69,70',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (7,17,'2020-03-30 17:00:00','锤子你好','','image','71,72,73,74,75,76',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (8,18,'2020-03-28 22:30:00','小鲤鱼历险记','','image','77,78,79,80',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (9,11,'2020-03-19 15:14:00','好久不见呀','','video','1',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (10,12,'2020-03-20 21:05:00','听说漂亮会传染 你靠我近一点','','video','2',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (11,13,'2020-03-21 20:40:00','换了一个冷酷无情的新发色','','video','3',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (12,14,'2020-01-21 22:15:00','男朋友和几十支口红同时掉进水里，你先救哪个色号？','','video','4',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (13,15,'2020-03-25 18:33:00','少一份矫情，多一份努力','','video','5',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (14,16,'2020-03-30 23:02:00','你们假期都在干什么？','','video','6',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (15,17,'2020-04-05 09:09:00','飒～','','video','7',1,NULL);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`) values (16,18,'2020-04-19 17:44:00','大噶猴，我是步惊云的弟弟步经风','','video','8',1,NULL);

/*Data for the table `dynamic_comment` */

/*Data for the table `dynamic_comment_like` */

/*Data for the table `dynamic_like` */

/*Data for the table `feedback` */

/*Data for the table `focus_notice` */

/*Data for the table `focuslist` */

/*Data for the table `gift` */

insert into `gift` (`id`, `price`, `name`) values('1','39','钻石鞋');
insert into `gift` (`id`, `price`, `name`) values('2','90','游轮');
insert into `gift` (`id`, `price`, `name`) values('3','77','炫酷跑车');
insert into `gift` (`id`, `price`, `name`) values('4','80','霸气摩托');
insert into `gift` (`id`, `price`, `name`) values('5','33','幸运飞碟');
insert into `gift` (`id`, `price`, `name`) values('6','12','桃花梦');
insert into `gift` (`id`, `price`, `name`) values('7','9','生日快乐');
insert into `gift` (`id`, `price`, `name`) values('8','55','求婚');
insert into `gift` (`id`, `price`, `name`) values('9','66','秋千少女');
insert into `gift` (`id`, `price`, `name`) values('10','88','魔法城堡');
insert into `gift` (`id`, `price`, `name`) values('11','33','浪漫依偎');
insert into `gift` (`id`, `price`, `name`) values('12','99','酷炫飞机');
insert into `gift` (`id`, `price`, `name`) values('13','88','节日烟花');
insert into `gift` (`id`, `price`, `name`) values('14','55','火箭腾飞');
insert into `gift` (`id`, `price`, `name`) values('15','29','皇冠');
insert into `gift` (`id`, `price`, `name`) values('16','88','欢快乐园');

/*Data for the table `gift_log` */

/*Data for the table `image` */

insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (1,0,'2020-08-05 13:43:12','avatar_0.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (2,1,'2020-08-05 13:43:12','avatar_1.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (3,2,'2020-08-05 13:43:12','avatar_2.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (4,3,'2020-08-05 13:43:12','avatar_3.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (5,4,'2020-08-05 13:43:12','avatar_4.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (6,5,'2020-08-05 13:43:13','avatar_5.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (7,6,'2020-08-05 13:43:13','avatar_6.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (8,7,'2020-08-05 13:43:13','avatar_7.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (9,8,'2020-08-05 13:43:13','avatar_8.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (10,9,'2020-08-05 13:43:13','avatar_9.png','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (11,11,'2020-08-05 13:43:13','avatar_11.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (12,12,'2020-08-05 13:43:13','avatar_12.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (13,13,'2020-08-05 13:43:13','avatar_13.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (14,14,'2020-08-05 13:43:13','avatar_14.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (15,15,'2020-08-05 13:43:13','avatar_15.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (16,16,'2020-08-05 13:43:13','avatar_16.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (17,17,'2020-08-05 13:43:13','avatar_17.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (18,18,'2020-08-05 13:43:13','avatar_18.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (19,11,'2020-08-05 13:43:13','photo_11_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (20,11,'2020-08-05 13:43:13','photo_11_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (21,11,'2020-08-05 13:43:13','photo_11_3.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (22,11,'2020-08-05 13:43:13','photo_11_4.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (23,11,'2020-08-05 13:43:13','photo_11_5.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (24,11,'2020-08-05 13:43:13','photo_11_6.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (25,12,'2020-08-05 13:43:13','photo_12_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (26,12,'2020-08-05 13:43:13','photo_12_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (27,13,'2020-08-05 13:43:13','photo_13_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (28,13,'2020-08-05 13:43:13','photo_13_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (29,14,'2020-08-05 13:43:13','photo_14_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (30,14,'2020-08-05 13:43:13','photo_14_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (31,14,'2020-08-05 13:43:13','photo_14_3.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (32,14,'2020-08-05 13:43:13','photo_14_4.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (33,14,'2020-08-05 13:43:13','photo_14_5.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (34,14,'2020-08-05 13:43:13','photo_14_6.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (35,15,'2020-08-05 13:43:13','photo_15_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (36,15,'2020-08-05 13:43:13','photo_15_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (37,16,'2020-08-05 13:43:13','photo_16_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (38,16,'2020-08-05 13:43:13','photo_16_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (39,16,'2020-08-05 13:43:14','photo_16_3.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (40,17,'2020-08-05 13:43:14','photo_17_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (41,17,'2020-08-05 13:43:14','photo_17_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (42,17,'2020-08-05 13:43:14','photo_17_3.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (43,18,'2020-08-05 13:43:14','photo_18_1.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (44,18,'2020-08-05 13:43:14','photo_18_2.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (45,18,'2020-08-05 13:43:14','photo_18_3.jpg','jpg','photo',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (46,11,'2020-08-05 13:43:14','dynamic_11_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (47,11,'2020-08-05 13:43:14','dynamic_11_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (48,11,'2020-08-05 13:43:14','dynamic_11_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (49,11,'2020-08-05 13:43:14','dynamic_11_4.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (50,11,'2020-08-05 13:43:14','dynamic_11_5.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (51,11,'2020-08-05 13:43:14','dynamic_11_6.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (52,12,'2020-08-05 13:43:14','dynamic_12_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (53,12,'2020-08-05 13:43:14','dynamic_12_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (54,12,'2020-08-05 13:43:14','dynamic_12_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (55,13,'2020-08-05 13:43:14','dynamic_13_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (56,13,'2020-08-05 13:43:14','dynamic_13_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (57,13,'2020-08-05 13:43:14','dynamic_13_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (58,14,'2020-08-05 13:43:14','dynamic_14_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (59,14,'2020-08-05 13:43:14','dynamic_14_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (60,14,'2020-08-05 13:43:14','dynamic_14_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (61,14,'2020-08-05 13:43:14','dynamic_14_4.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (62,14,'2020-08-05 13:43:14','dynamic_14_5.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (63,14,'2020-08-05 13:43:14','dynamic_14_6.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (64,15,'2020-08-05 13:43:14','dynamic_15_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (65,15,'2020-08-05 13:43:14','dynamic_15_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (66,15,'2020-08-05 13:43:14','dynamic_15_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (67,16,'2020-08-05 13:43:14','dynamic_16_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (68,16,'2020-08-05 13:43:14','dynamic_16_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (69,16,'2020-08-05 13:43:14','dynamic_16_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (70,16,'2020-08-05 13:43:14','dynamic_16_4.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (71,16,'2020-08-05 13:43:14','dynamic_17_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (72,16,'2020-08-05 13:43:14','dynamic_17_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (73,16,'2020-08-05 13:43:15','dynamic_17_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (74,16,'2020-08-05 13:43:15','dynamic_17_4.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (75,16,'2020-08-05 13:43:15','dynamic_17_5.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (76,16,'2020-08-05 13:43:15','dynamic_17_6.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (77,18,'2020-08-05 13:43:15','dynamic_18_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (78,18,'2020-08-05 13:43:15','dynamic_18_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (79,18,'2020-08-05 13:43:15','dynamic_18_3.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (80,18,'2020-08-05 13:43:15','dynamic_18_4.jpg','jpg','dynamic',1,NULL);

/*Data for the table `match_log` */

/*Data for the table `pay_config` */

insert  into `pay_config`(`id`,`money`,`coins`,`appid`) values (1,12,120,'1000000000000001');
insert  into `pay_config`(`id`,`money`,`coins`,`appid`) values (2,50,500,'1000000000000002');
insert  into `pay_config`(`id`,`money`,`coins`,`appid`) values (3,188,1880,'1000000000000003');

/*Data for the table `pay_order` */

/*Data for the table `photolist` */

insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (1,11,'2020-08-05 13:43:15','19,20,21,22,23,24',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (2,12,'2020-08-05 13:43:15','25,26',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (3,13,'2020-08-05 13:43:15','27,28',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (4,14,'2020-08-05 13:43:15','29,30,31,32,33,34',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (5,15,'2020-08-05 13:43:15','35,36',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (6,16,'2020-08-05 13:43:15','37,38,39',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (7,17,'2020-08-05 13:43:16','40,41,42',1,NULL);
insert  into `photolist`(`id`,`user_id`,`post_time`,`photolist`,`is_audit`,`audit_time`) values (8,18,'2020-08-05 13:43:16','43,44,45',1,NULL);

/*Data for the table `room` */

insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`) values (1,12,1,'@TGS#1MFHZBTGR','进来聊天吧',0,1,'2020-07-14 14:30:00',NULL);
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`) values (2,13,0,'@TGS#1QKIZBTG3','我是个敲可爱',0,1,'2020-07-14 14:20:00',NULL);
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`) values (3,14,0,'@TGS#167IZBTGL','今日份元气已加满',0,1,'2020-07-14 14:10:00',NULL);
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`) values (4,15,1,'@TGS#1BHJZBTGK','圆圆的脸蛋',0,1,'2020-07-14 14:00:00',NULL);

/*Data for the table `room_seat` */

/*Data for the table `room_user` */

/*Data for the table `user` */

insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (1,'19800000001','2020-08-05 11:24:51','系统提醒',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',2,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (2,'19800000002','2020-08-05 11:24:51','反馈与帮助',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',3,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (3,'19800000003','2020-08-05 11:24:51','预留',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (4,'19800000004','2020-08-05 11:24:51','预留',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (5,'19800000005','2020-08-05 11:24:51','预留',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (6,'19800000006','2020-08-05 11:24:51','预留',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (7,'19800000007','2020-08-05 11:24:51','预留',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (8,'19800000008','2020-08-05 11:24:51','预留',1,'1900-01-01','','2020-08-05 11:24:51','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (9,'19800000009','2020-08-05 11:24:52','预留',1,'1900-01-01','','2020-08-05 11:24:52','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (10,'19800000010','2020-08-05 11:24:52','预留',1,'1900-01-01','','2020-08-05 11:24:52','0.0.0.0',1,0,0,'','','','',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (11,'19900000001','2020-08-05 11:24:52','IMMune',0,'1998-07-15','','2020-08-05 11:24:52','0.0.0.0',11,1,0,'热爱生活，欣赏身边美好的食物','','找旅游伴侣','潜水,音乐',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (12,'19900000002','2020-08-05 11:24:52','莹莹子',0,'1994-09-18','','2020-08-05 11:24:52','0.0.0.0',12,2,0,'你会遇见更好的自己','单身','交朋友','唱歌,美食',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (13,'19900000003','2020-08-05 11:24:52','一只丸子兔',0,'1990-06-12','','2020-08-05 11:24:52','0.0.0.0',13,3,0,'没什么不同','','约会','沉迷电竞',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (14,'19900000004','2020-08-05 11:24:52','Milk.疯丫头',0,'1991-03-03','','2020-08-05 11:24:52','0.0.0.0',14,4,0,'变成了星星','','聊天','写作',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (15,'19900000005','2020-08-05 11:24:52','子言呀',1,'1997-02-09','','2020-08-05 11:24:52','0.0.0.0',15,5,0,'天外来物','单身','一起玩游戏','玩手工,音乐',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (16,'19900000006','2020-08-05 11:24:52','Alisa孟',1,'1992-04-05','','2020-08-05 11:24:52','0.0.0.0',16,6,0,'眼里只有你','恋爱','聊天','综艺梗',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (17,'19900000007','2020-08-05 11:24:52','pink ink',1,'1998-06-23','','2020-08-05 11:24:52','0.0.0.0',17,7,0,'饼干爱远杰','','约会','美食,交朋友',0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`) values (18,'19900000008','2020-08-05 11:24:52','Tattoo奕',1,'1995-03-24','','2020-08-05 11:24:52','0.0.0.0',18,8,0,'有你的moonlight','恋爱','一起玩游戏','动漫,科技',0,0);

/*Data for the table `video` */

insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (1,11,'2020-08-05 13:43:15','dynamic_11.mp4','mp4','dynamic_11.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (2,12,'2020-08-05 13:43:15','dynamic_12.mp4','mp4','dynamic_12.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (3,13,'2020-08-05 13:43:15','dynamic_13.mp4','mp4','dynamic_13.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (4,14,'2020-08-05 13:43:15','dynamic_14.mp4','mp4','dynamic_14.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (5,15,'2020-08-05 13:43:15','dynamic_15.mp4','mp4','dynamic_15.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (6,16,'2020-08-05 13:43:15','dynamic_16.mp4','mp4','dynamic_16.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (7,17,'2020-08-05 13:43:15','dynamic_17.mp4','mp4','dynamic_17.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (8,18,'2020-08-05 13:43:15','dynamic_18.mp4','mp4','dynamic_18.jpg','jpg',0,'dynamic',1,NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
