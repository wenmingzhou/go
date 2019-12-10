-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2019-12-10 14:04:12
-- 服务器版本： 8.0.12
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `ego`
--

-- --------------------------------------------------------

--
-- 表的结构 `tb_item`
--

CREATE TABLE `tb_item` (
  `id` bigint(30) NOT NULL,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品标题',
  `sell_point` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品卖点',
  `price` int(5) NOT NULL COMMENT '商品价格，单位为:分',
  `num` int(5) NOT NULL COMMENT '库存数量',
  `barcode` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品条形码',
  `image` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品图片',
  `cid` int(10) NOT NULL COMMENT '所属类目',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '(1->正常 2->下架 3->删除)',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `tb_item`
--

INSERT INTO `tb_item` (`id`, `title`, `sell_point`, `price`, `num`, `barcode`, `image`, `cid`, `status`, `created`, `updated`) VALUES
(1575425299, '标题1', '卖点2777', 999, 11, '', 'http://localhost:9999/static/images/1575425291427893200992.jpg,http://localhost:9999/static/images/1575425291482896300646.jpg', 3, 1, '2019-12-04 10:08:19', '2019-12-05 10:03:17'),
(93141575425399, '12', '233', 0, 0, '', '', 1, 1, '2019-12-04 10:09:59', '2019-12-04 10:09:59'),
(86401575425409, '12', '233', 0, 0, '', '', 1, 1, '2019-12-04 10:10:09', '2019-12-04 10:10:09'),
(22561575425942, '11', '22', 33, 44, '', '', 3, 1, '2019-12-04 10:19:02', '2019-12-04 10:19:02'),
(44711575428055, '11', '222', 99, 99, '', 'http://localhost:9999/static/images/1575428040736144600510.jpg,http://localhost:9999/static/images/1575428040817149300596.jpg', 3, 1, '2019-12-04 10:54:15', '2019-12-04 10:54:15');

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_cat`
--

CREATE TABLE `tb_item_cat` (
  `id` int(10) NOT NULL,
  `parent_id` int(10) NOT NULL COMMENT '父类目ID=0时，代表的是一级类目',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '类目名称',
  `status` int(1) NOT NULL DEFAULT '1' COMMENT '1(正常)2(删除)',
  `sort_order` int(5) NOT NULL COMMENT '排序序号',
  `is_parent` tinyint(1) NOT NULL COMMENT '该类目是否为父类目1为true,0为false',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `tb_item_cat`
--

INSERT INTO `tb_item_cat` (`id`, `parent_id`, `name`, `status`, `sort_order`, `is_parent`, `created`, `updated`) VALUES
(1, 0, '图书，影像，电子书刊', 1, 1, 1, '2019-11-29 13:41:56', '2019-11-29 13:42:00'),
(2, 0, '家用电器', 1, 1, 1, '2019-11-29 13:41:56', '2019-11-29 13:42:00'),
(3, 1, '电子书刊', 1, 1, 0, '2019-11-29 13:41:56', '2019-11-29 13:42:00'),
(4, 1, '音像', 1, 1, 0, '2019-11-29 13:41:56', '2019-11-29 13:42:00');

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_desc`
--

CREATE TABLE `tb_item_desc` (
  `item_id` bigint(20) NOT NULL COMMENT '鍟嗗搧id',
  `item_desc` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品描述',
  `created` datetime NOT NULL COMMENT '创建时间',
  `updated` datetime NOT NULL COMMENT '更新时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `tb_item_desc`
--

INSERT INTO `tb_item_desc` (`item_id`, `item_desc`, `created`, `updated`) VALUES
(44711575428055, '描述描述', '2019-12-04 10:54:15', '2019-12-04 10:54:15'),
(1575425299, '112777', '2019-12-04 10:55:15', '2019-12-05 10:03:17');

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_param`
--

CREATE TABLE `tb_item_param` (
  `id` bigint(20) NOT NULL,
  `item_cat_id` bigint(20) NOT NULL COMMENT '商品类目id',
  `param_data` text NOT NULL COMMENT '参数数据，json格式',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `tb_item_param`
--

INSERT INTO `tb_item_param` (`id`, `item_cat_id`, `param_data`, `created`, `updated`) VALUES
(2, 1, '122', '2019-12-05 00:00:00', '2019-12-05 00:00:00');

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_param_item`
--

CREATE TABLE `tb_item_param_item` (
  `id` bigint(20) NOT NULL,
  `item_id` int(11) NOT NULL COMMENT '商品id',
  `param_data` int(11) NOT NULL COMMENT '参数数据，json格式',
  `created` int(11) NOT NULL,
  `updated` int(11) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_user`
--

CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(32) NOT NULL,
  `phone` varchar(30) NOT NULL,
  `email` varchar(30) NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `tb_user`
--

INSERT INTO `tb_user` (`id`, `username`, `password`, `phone`, `email`, `created`, `updated`) VALUES
(1, 'root', 'root', '13770707526', '524797132@qq.com', '2019-11-25 07:00:00', '2019-11-25 09:00:00');

--
-- 转储表的索引
--

--
-- 表的索引 `tb_item`
--
ALTER TABLE `tb_item`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tb_item_cat`
--
ALTER TABLE `tb_item_cat`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tb_item_param`
--
ALTER TABLE `tb_item_param`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tb_item_param_item`
--
ALTER TABLE `tb_item_param_item`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tb_user`
--
ALTER TABLE `tb_user`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `tb_item_cat`
--
ALTER TABLE `tb_item_cat`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `tb_item_param`
--
ALTER TABLE `tb_item_param`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `tb_item_param_item`
--
ALTER TABLE `tb_item_param_item`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `tb_user`
--
ALTER TABLE `tb_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
