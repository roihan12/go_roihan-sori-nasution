date=$(date '+ %a %b %T WIB %Y')


mkdir "$1 at $date"
cd "$1 at $date"
mkdir about_me my_friends my_system_info
cd about_me
mkdir personal professional
cd personal
echo "https://www.facebook.com/$2" > facebook.txt
cd ..
cd professional
echo "https://www.linkedin.com/in/$3" > linkedin.txt
cd ..
cd ..

cd my_friends
echo $(curl -XGET 'https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt') > list_of_my_friends.txt
cd ..
                                                                                                      
cd my_system_info
echo 'My username:' $(whoami) > about_this_laptop.txt
echo 'With host: ' $(uname -a) >> about_this_laptop.txt
echo $(ping google.com) > internet_connection.txt

