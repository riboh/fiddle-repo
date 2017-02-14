// compile error用ファイル
public class StaticError{
  public static void main(String args[]){
    final String COLOR = "blue";
    COLOR = "red";
    // other code
    System.out.println(COLOR);
  }
}
