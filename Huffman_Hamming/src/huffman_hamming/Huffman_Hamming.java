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
        // TODO code application logic here 
       //"com.sun.java.swing.plaf.nimbus.NimbusLookAndFeel"
    
       System.setProperty("apple.laf.useScreenMenuBar", "true");
       System.setProperty("com.apple.mrj.application.apple.menu.about.name", "WikiTeX");
       try {
           UIManager.setLookAndFeel(UIManager.getSystemLookAndFeelClassName());
       } catch (Exception e) {
           e.printStackTrace(System.out);
       }
       javax.swing.JFrame ventana = new VentanaPrincipal();
         
       ventana.setVisible(true);
    }
}
