package kata


func BinToDec(bin string) int {
      
  
    //Bin is a string and return integer 
    // we have to convert string to integer by usinga  manaul code
  
    
    // we have to convert Bin string to Integer 
    // and write the name of bin to decmaail . 
    //
    // return the intger 
      
  // convert string to intger "10101" = 10101
    num:= 0 
    
  
 
    for i:= 0; i  <= len(bin)-1 ; i++{
      num = num*2+ (int(bin[i]-48))
    }
    
    
  return num
}


//power 2**4 = 16 that's how the power are made bro. 

