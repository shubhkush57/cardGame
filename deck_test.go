 package main
import "testing"
 func TestNewDeck(t *testing.T){
	 d := newDec()
	 if( len(d) != 16){
		t.Errorf("Expected len of 16 but got %d",len(d))
	 }
	 if(d[0] != "Ace of Spades"){
		t.Errorf("Expeced first Ace of Spades but got %v",d[0])
	 }
	 
 }