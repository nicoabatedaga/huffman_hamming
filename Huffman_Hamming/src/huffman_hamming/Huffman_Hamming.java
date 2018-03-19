/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package huffman_hamming;
import javax.swing.*;          
import java.awt.*;
import java.awt.event.*;
import javax.swing.plaf.metal.*;

/**
 *
 * @author Joaquin
 */
public class Huffman_Hamming { 
    /**
     * @param args the command line arguments
     */
    
 
    
    public static void main(String[] args) {  
       try {
           UIManager.setLookAndFeel(UIManager.getSystemLookAndFeelClassName());
       } catch (Exception e) {
           e.printStackTrace(System.out);
       }
       javax.swing.JFrame ventana = new VentanaPrincipal();   
       /*Image image = Toolkit.getDefaultToolkit().getImage(".\\resource\\unsl.ico"); 
       
       ventana.setIconImage(image);        */ 
       ventana.setVisible(true);
    }
}
