BEGIN{ fn = "part1.txt"; n = 1 }
{
   print > fn
   if (substr($0,1,2) == "}") {
       close (fn)
       n++
       fn = "part" n ".txt"
   }
}

