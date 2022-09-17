ping =ping -c3 Google.com

mkdir "tegar at $(date '+ %a %b %T WIB %Y')"
cd "tegar at $(date '+ %a %b %T WIB %Y')"
mkdir about_me my_friends my_system_info
cd about_me
mkdir personal professional
cd personal
echo "https://www.facebook.com/tegarimansyahfb" > facebook.txt
cd ..
cd professional
echo "https://www.linkedin.com/in/tegarimansyahlinkedin" > linkedin.txt
cd ..
cd ..

cd my_friends
curl "https://gist.github.com/tegarimansyah/e91f335753ab2c7fb12815779677e914#file-list-of-my-friends-txt" > list- of-my-friends.txt
cd ..

cd my_system_info
echo "My username: $USER \nWith host: uname -a" >> about_this_laptop.txt
echo $ping >> internet_connection.txt
