#! /bin/bash  # employ bash shell

echo -e "\033[32m
-----------------------------------------------------------------------------------------------------------------------------------------------------
		            一键申请https 证书
-----------------------------------------------------------------------------------------------------------------------------------------------------
	请选择您的域名提供商(请先安装acme.sh,复制路径回车:curl https://get.acme.sh | sh)
	1.CloudFlare
	2.GoDaddy.com
	3.阿里云域API
	4.DNSPod.cn

\033[0m"
read num
if [ $num -eq 1 ]; then
	echo "please input your CF_key:"
	read key1
	echo "your CF_Email:" 
	read key2
	export CF_key="$key1"
	export CF_Email="$key2"
	dns_class="dns_cf"
elif [ $num -eq 2 ]; then
	echo "please input your GD_Key:"
	read key1
	echo "your GD_Secret:" 
	read key2
      export GD_Key="$key1"
	export GD_Secret="$key2"
	dns_class="dns_gd"
elif [ $num -eq 2 ]; then
	echo "please input your Ali_Key:"
	read key1
	echo "your Ali_Secret:" 
	read key2
      export Ali_Key="$key1"
	export Ali_Secret="$key2"
	dns_class="dns_ali"

 else
	echo "输入有误，请重新输入"
	bash myssl.sh
fi

echo "your server_name 你的域名 （如： gojw.xyz） :" 
read hostName
echo "your key&cer of location 你的证书和密钥存储位置 (如:  /usr/local  ):" 
read location

echo "正在申请，请耐心等待........"
~/.acme.sh/acme.sh   --issue -d $hostName  -d *.$hostName --dns $dns_class

~/.acme.sh/acme.sh  --installcert  -d  $hostName   \
        --key-file   $location/$hostName.key \
        --fullchain-file $location/fullchain.cer \
