/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package huffman_hamming;

import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.ItemEvent;
import java.awt.event.ItemListener;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.util.Properties;
import java.util.Vector;
import java.util.prefs.Preferences;
import javax.swing.Action;
import javax.swing.ButtonGroup;
import javax.swing.DefaultComboBoxModel;
import javax.swing.JComboBox;
import javax.swing.JFileChooser;
import javax.swing.JFrame;
import javax.swing.JMenuItem;
import javax.swing.JOptionPane;
import javax.swing.JPanel;
import javax.swing.JPopupMenu;
import javax.swing.JRadioButton;
import javax.swing.UIManager;
import javax.swing.filechooser.FileFilter;
import javax.swing.filechooser.FileView;

/**
 *
 * @author Joaquin
 */
public class VentanaPrincipal extends javax.swing.JFrame {
    /**C:\Users\Joaquin\go\src\huffman_hamming
     * Creates new form VentanaPrincipal
     */
    int codificacion =512;
    String editorPreferido ="notepad";
    private String ejecutablePath="C:\\Users\\Joaquin\\go\\src\\huffman_hamming\\huffman_hamming.exe"; 
    public void cargarPreferencias(){
        Properties prop = new Properties();
        InputStream input = null;
        try {
            input = new FileInputStream("./preferencias");
            prop.load(input);
            editorPreferido=(prop.getProperty("editor"));
            ejecutablePath=(prop.getProperty("ejecutable")); 
        } catch (IOException ex) {
            ex.printStackTrace();
        } finally {
            if (input != null) {
                    try {
                            input.close();
                    } catch (IOException e) {
                            e.printStackTrace();
                    }
            }
        }
    }
    public void restablercerPreferencias(){
        try{    
            String EDITOR = "notepad";
            String EJECUTABLE = "C:\\Users\\Joaquin\\go\\src\\huffman_hamming\\huffman_hamming.exe";

             JPanel panel = new JPanel();
            JFileChooser fc = new JFileChooser();
            panel.add(fc);
            fc.setDialogTitle("Seleccione ejecutable programa");
            fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
            fc.setFileFilter(new FileFilter() {
                @Override
                public boolean accept(File f) {
                    if ( getExtension(f.getName()).equals("exe")){
                        return true;
                    }else{
                        if (f.isDirectory()){
                        return true;
                        }
                    }
                    return false;
                }
                @Override
                public String getDescription() {
                    return "Seleccionar archivos .exe como editor de texto.";
                }
            });
            fc.setApproveButtonText("Ejecutable"); 
            panel.setVisible(true);
            int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
            if (seleccion == JFileChooser.APPROVE_OPTION){
                File fileOut = fc.getSelectedFile();
                 if (fileOut != null){  
                     EJECUTABLE=fileOut.getAbsolutePath();
                }
            }else{
                System.out.print("ERROR: NO SELECCIONO PATH");
                System.exit(0);
            }
             //create a properties file
             Properties props = new Properties();
             props.setProperty("editor", EDITOR);
             props.setProperty("ejecutable", EJECUTABLE);
             File f = new File("./preferencias");
             OutputStream out = new FileOutputStream( f );
             //If you wish to make some comments 
             props.store(out, "Preferencias");
         }
          catch (Exception e ) {
              e.printStackTrace();
          }
    }
   public void guardarPreferencias(){
   try{    
        String EDITOR = editorPreferido;
        String EJECUTABLE = ejecutablePath;
        //create a properties file
        Properties props = new Properties();
        props.setProperty("editor", EDITOR);
        props.setProperty("ejecutable", EJECUTABLE);
        File f = new File("./preferencias");
        OutputStream out = new FileOutputStream( f );
        //If you wish to make some comments 
        props.store(out, "Preferencias");
   }
    catch (Exception e ) {
        e.printStackTrace();
    }
   }

    public VentanaPrincipal() {
        
        initComponents();  
        this.botonComprimir.setEnabled(true);
        this.botonComprobar.setEnabled(false);           
        this.botonCorregir.setEnabled(false);
        this.botonDañar.setEnabled(false);
        this.botonDescomprimir.setEnabled(false);
        this.botonDesproteger.setEnabled(false);
        this.botonProteger.setEnabled(true);
        this.botonVerArbol.setEnabled(false);
        this.botonComprimir.setVisible(true);
        this.botonComprobar.setVisible(false);           
        this.botonCorregir.setVisible(false);
        this.botonDañar.setVisible(false);
        this.botonDescomprimir.setVisible(false);
        this.botonDesproteger.setVisible(false);
        this.botonProteger.setVisible(true);
        this.botonVerArbol.setVisible(false);
        File pref= new File("./preferencias");
        if (!pref.exists() ){
            restablercerPreferencias();
        }
        cargarPreferencias();

        //this.seletorArchivosIn.remove(this.seletorArchivosIn.get);
        ButtonGroup bG=new ButtonGroup();

        JRadioButton jr512=new JRadioButton("512",true);
        bG.add(jr512);
        jr512.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                codificacion = 512;
            }
        });
        JRadioButton jr1024=new JRadioButton("1024",false);
        bG.add(jr1024);
        jr1024.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                codificacion = 1024;
            }
        });
        JRadioButton jr12048=new JRadioButton("2048",false);
        bG.add(jr12048);
        jr12048.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                codificacion = 2048;
            }
        });
        JMenuItem item512=new JMenuItem("512");
        JMenuItem item1024=new JMenuItem("1024");
        JMenuItem item2048=new JMenuItem("2048");

        item512.add("512",jr512); 
        item1024.add("1024",jr1024); 
        item2048.add("2048",jr12048); 

        
       JMenuItem item =this.menuCodificacion.add( item512);    
       this.menuCodificacion.add( item1024);
       this.menuCodificacion.add( item2048);

    }

    /**
     * This method is called from within the constructor to initialize the form.
     * WARNING: Do NOT modify this code. The content of this method is always
     * regenerated by the Form Editor.
     */
    @SuppressWarnings("unchecked")
    // <editor-fold defaultstate="collapsed" desc="Generated Code">//GEN-BEGIN:initComponents
    private void initComponents() {

        jPanel1 = new javax.swing.JPanel();
        seletorArchivosIn = new javax.swing.JFileChooser();
        jPanel2 = new javax.swing.JPanel();
        botonComprimir = new javax.swing.JButton();
        botonDescomprimir = new javax.swing.JButton();
        botonVerArbol = new javax.swing.JButton();
        botonProteger = new javax.swing.JButton();
        botonDesproteger = new javax.swing.JButton();
        botonComprobar = new javax.swing.JButton();
        botonDañar = new javax.swing.JButton();
        botonCorregir = new javax.swing.JButton();
        botonVerMatrices = new javax.swing.JButton();
        BarraMenu = new javax.swing.JMenuBar();
        jMenu1 = new javax.swing.JMenu();
        botonEjecutable = new javax.swing.JMenuItem();
        botonEditor = new javax.swing.JMenuItem();
        menuCodificacion = new javax.swing.JMenu();
        verConfiguracion = new javax.swing.JMenuItem();
        botonRestablecer = new javax.swing.JMenuItem();
        jMenu2 = new javax.swing.JMenu();

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        setResizable(false);
        addWindowListener(new java.awt.event.WindowAdapter() {
            public void windowClosed(java.awt.event.WindowEvent evt) {
                formWindowClosed(evt);
            }
            public void windowClosing(java.awt.event.WindowEvent evt) {
                formWindowClosing(evt);
            }
        });

        seletorArchivosIn.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                seletorArchivosInActionPerformed(evt);
            }
        });
        seletorArchivosIn.addPropertyChangeListener(new java.beans.PropertyChangeListener() {
            public void propertyChange(java.beans.PropertyChangeEvent evt) {
                seletorArchivosInPropertyChange(evt);
            }
        });

        botonComprimir.setText("Comprimir");
        botonComprimir.setMaximumSize(new java.awt.Dimension(100, 35));
        botonComprimir.setMinimumSize(new java.awt.Dimension(100, 35));
        botonComprimir.setPreferredSize(new java.awt.Dimension(100, 35));
        botonComprimir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonComprimirActionPerformed(evt);
            }
        });

        botonDescomprimir.setText("Descomprimir");
        botonDescomprimir.setMaximumSize(new java.awt.Dimension(100, 35));
        botonDescomprimir.setMinimumSize(new java.awt.Dimension(100, 35));
        botonDescomprimir.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDescomprimir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDescomprimirActionPerformed(evt);
            }
        });

        botonVerArbol.setText("Ver Arbol");
        botonVerArbol.setMaximumSize(new java.awt.Dimension(100, 35));
        botonVerArbol.setMinimumSize(new java.awt.Dimension(100, 35));
        botonVerArbol.setPreferredSize(new java.awt.Dimension(100, 35));
        botonVerArbol.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonVerArbolActionPerformed(evt);
            }
        });

        botonProteger.setText("Proteger");
        botonProteger.setMaximumSize(new java.awt.Dimension(100, 35));
        botonProteger.setMinimumSize(new java.awt.Dimension(100, 35));
        botonProteger.setPreferredSize(new java.awt.Dimension(100, 35));
        botonProteger.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonProtegerActionPerformed(evt);
            }
        });

        botonDesproteger.setText("Desproteger");
        botonDesproteger.setMaximumSize(new java.awt.Dimension(100, 35));
        botonDesproteger.setMinimumSize(new java.awt.Dimension(100, 35));
        botonDesproteger.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDesproteger.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDesprotegerActionPerformed(evt);
            }
        });

        botonComprobar.setText("Comprobar");
        botonComprobar.setMaximumSize(new java.awt.Dimension(100, 35));
        botonComprobar.setMinimumSize(new java.awt.Dimension(100, 35));
        botonComprobar.setPreferredSize(new java.awt.Dimension(100, 35));
        botonComprobar.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonComprobarActionPerformed(evt);
            }
        });

        botonDañar.setText("Dañar");
        botonDañar.setMaximumSize(new java.awt.Dimension(100, 35));
        botonDañar.setMinimumSize(new java.awt.Dimension(100, 35));
        botonDañar.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDañar.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDañarActionPerformed(evt);
            }
        });

        botonCorregir.setText("Corregir");
        botonCorregir.setMaximumSize(new java.awt.Dimension(100, 35));
        botonCorregir.setMinimumSize(new java.awt.Dimension(100, 35));
        botonCorregir.setPreferredSize(new java.awt.Dimension(100, 35));
        botonCorregir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonCorregirActionPerformed(evt);
            }
        });

        botonVerMatrices.setText("Ver Matrices");
        botonVerMatrices.setMaximumSize(new java.awt.Dimension(100, 35));
        botonVerMatrices.setMinimumSize(new java.awt.Dimension(100, 35));
        botonVerMatrices.setPreferredSize(new java.awt.Dimension(100, 35));
        botonVerMatrices.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonVerMatricesActionPerformed(evt);
            }
        });

        javax.swing.GroupLayout jPanel2Layout = new javax.swing.GroupLayout(jPanel2);
        jPanel2.setLayout(jPanel2Layout);
        jPanel2Layout.setHorizontalGroup(
            jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel2Layout.createSequentialGroup()
                .addContainerGap()
                .addComponent(botonComprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonDescomprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonVerArbol, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonProteger, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonDesproteger, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonComprobar, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonDañar, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonCorregir, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonVerMatrices, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );
        jPanel2Layout.setVerticalGroup(
            jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel2Layout.createSequentialGroup()
                .addContainerGap()
                .addGroup(jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                    .addComponent(botonDañar, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonComprobar, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonDesproteger, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonProteger, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonVerArbol, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonDescomprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonComprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonCorregir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(botonVerMatrices, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE))
                .addContainerGap())
        );

        javax.swing.GroupLayout jPanel1Layout = new javax.swing.GroupLayout(jPanel1);
        jPanel1.setLayout(jPanel1Layout);
        jPanel1Layout.setHorizontalGroup(
            jPanel1Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jPanel2, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
            .addGroup(javax.swing.GroupLayout.Alignment.TRAILING, jPanel1Layout.createSequentialGroup()
                .addComponent(seletorArchivosIn, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                .addContainerGap())
        );
        jPanel1Layout.setVerticalGroup(
            jPanel1Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel1Layout.createSequentialGroup()
                .addComponent(jPanel2, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(seletorArchivosIn, javax.swing.GroupLayout.PREFERRED_SIZE, 371, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );

        jMenu1.setText("Preferencias");

        botonEjecutable.setText("Ejecutable");
        botonEjecutable.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonEjecutableActionPerformed(evt);
            }
        });
        jMenu1.add(botonEjecutable);

        botonEditor.setText("Editor de texto");
        botonEditor.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonEditorActionPerformed(evt);
            }
        });
        jMenu1.add(botonEditor);

        menuCodificacion.setText("Codificación");
        jMenu1.add(menuCodificacion);

        verConfiguracion.setText("Ver configuración");
        verConfiguracion.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                verConfiguracionActionPerformed(evt);
            }
        });
        jMenu1.add(verConfiguracion);

        botonRestablecer.setText("Restabler configuración");
        botonRestablecer.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonRestablecerActionPerformed(evt);
            }
        });
        jMenu1.add(botonRestablecer);

        BarraMenu.add(jMenu1);

        jMenu2.setText("Información");
        BarraMenu.add(jMenu2);

        setJMenuBar(BarraMenu);

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(getContentPane());
        getContentPane().setLayout(layout);
        layout.setHorizontalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jPanel1, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
        );
        layout.setVerticalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(layout.createSequentialGroup()
                .addComponent(jPanel1, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                .addContainerGap())
        );

        pack();
    }// </editor-fold>//GEN-END:initComponents

    private void botonVerArbolActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonVerArbolActionPerformed
        // TODO add your handling code here:
                System.out.println("Ver arbol");

    }//GEN-LAST:event_botonVerArbolActionPerformed

    private void botonProtegerActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonProtegerActionPerformed
        File fileIn = this.seletorArchivosIn.getSelectedFile();

        String pathOut = "";   
        JPanel panel = new JPanel();
        JFileChooser fc = new JFileChooser();
        panel.add(fc);
        fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
        fc.setApproveButtonText("Proteger");
        panel.setVisible(true);
        int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
        if (seleccion == JFileChooser.APPROVE_OPTION){
            File fileOut = fc.getSelectedFile();
             if (fileOut != null){
                System.out.println(ejecutablePath+" -op=p -in="+fileIn+" -out="+fileOut+".ham -cod="+codificacion);
                 ejecutar(ejecutablePath+" -op=p -in="+fileIn+" -out="+fileOut+".ham -cod="+codificacion);
                 seletorArchivosIn.updateUI();
            }
        }
    }//GEN-LAST:event_botonProtegerActionPerformed

    
   public String ejecutar(String comando){
    StringBuffer  resultado =new StringBuffer(); 
    try{
        Process p=Runtime.getRuntime().exec(comando);
        p.waitFor();
        BufferedReader reader =
        new BufferedReader(new InputStreamReader(p.getInputStream()));
        String line = "";
        while ((line = reader.readLine())!= null) {
            resultado.append(line + "\n");
        }
    } catch (Exception e) {
        e.printStackTrace();
    }
    return resultado.toString();

   }
    private void seletorArchivosInPropertyChange(java.beans.PropertyChangeEvent evt) {//GEN-FIRST:event_seletorArchivosInPropertyChange
    if(JFileChooser.SELECTED_FILE_CHANGED_PROPERTY.equals(evt.getPropertyName())){
        File file =this.seletorArchivosIn.getSelectedFile();
        if (file != null){
            String extension =getExtension(file.toString());
            switch (extension){
                case "ham":
                    this.botonComprimir.setEnabled(false);
                    this.botonComprobar.setEnabled(true);
                    this.botonCorregir.setEnabled(true);
                    this.botonDañar.setEnabled(true);
                    this.botonDescomprimir.setEnabled(false);
                    this.botonDesproteger.setEnabled(true);
                    this.botonProteger.setEnabled(false);
                    this.botonVerArbol.setEnabled(false);
                    this.botonComprimir.setVisible(false);
                    this.botonComprobar.setVisible(true);
                    this.botonCorregir.setVisible(true);
                    this.botonDañar.setVisible(true);
                    this.botonDescomprimir.setVisible(false);
                    this.botonDesproteger.setVisible(true);
                    this.botonProteger.setVisible(false);
                    this.botonVerArbol.setVisible(false);
                    break;
                case "huf":
                    this.botonComprimir.setEnabled(false);
                    this.botonComprobar.setEnabled(false);
                    this.botonCorregir.setEnabled(false);
                    this.botonDañar.setEnabled(false);
                    this.botonDescomprimir.setEnabled(true);
                    this.botonDesproteger.setEnabled(false);
                    this.botonProteger.setEnabled(true);
                    this.botonVerArbol.setEnabled(true);
                    this.botonComprimir.setVisible(false);
                    this.botonComprobar.setVisible(false);
                    this.botonCorregir.setVisible(false);
                    this.botonDañar.setVisible(false);
                    this.botonDescomprimir.setVisible(true);
                    this.botonDesproteger.setVisible(false);
                    this.botonProteger.setVisible(true);
                    this.botonVerArbol.setVisible(true);
                    break;
                default:
                    this.botonComprimir.setEnabled(true);
                    this.botonComprobar.setEnabled(false);           
                    this.botonCorregir.setEnabled(false);
                    this.botonDañar.setEnabled(false);
                    this.botonDescomprimir.setEnabled(false);
                    this.botonDesproteger.setEnabled(false);
                    this.botonProteger.setEnabled(true);
                    this.botonVerArbol.setEnabled(false);
                    this.botonComprimir.setVisible(true);
                    this.botonComprobar.setVisible(false);           
                    this.botonCorregir.setVisible(false);
                    this.botonDañar.setVisible(false);
                    this.botonDescomprimir.setVisible(false);
                    this.botonDesproteger.setVisible(false);
                    this.botonProteger.setVisible(true);
                    this.botonVerArbol.setVisible(false);
                        
                    break;
                            
            
            
            }
            
        }
    }       
        
// TODO add your handling code here:
    }//GEN-LAST:event_seletorArchivosInPropertyChange

    private void botonComprimirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonComprimirActionPerformed
        // TODO add your handling code here:
        System.out.println("Comprimir");
    }//GEN-LAST:event_botonComprimirActionPerformed

    private void botonDescomprimirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonDescomprimirActionPerformed
        // TODO add your handling code here:
        System.out.println("Descomprimir");

    }//GEN-LAST:event_botonDescomprimirActionPerformed

    private void botonDesprotegerActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonDesprotegerActionPerformed
        // TODO add your handling code here:
        File fileIn = this.seletorArchivosIn.getSelectedFile();

        String pathOut = "";   
        JPanel panel = new JPanel();
        JFileChooser fc = new JFileChooser();
        panel.add(fc);
        fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
        fc.setApproveButtonText("Desproteger");
        panel.setVisible(true);
        int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
        if (seleccion == JFileChooser.APPROVE_OPTION){
            File fileOut = fc.getSelectedFile();
             if (fileOut != null){
                System.out.println(ejecutablePath+" -op=dp -in="+fileIn+" -out="+fileOut);
                String resultado=ejecutar(ejecutablePath+" -op=dp -in="+fileIn+" -out="+fileOut);
                if (resultado.length() >1){
                            JOptionPane.showMessageDialog(new JFrame(),resultado,"Desproteger error", JOptionPane.ERROR_MESSAGE);
                        }
                seletorArchivosIn.updateUI();
            }
        }
    }//GEN-LAST:event_botonDesprotegerActionPerformed

    private void botonComprobarActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonComprobarActionPerformed
        // TODO add your handling code here:
        File fileIn = this.seletorArchivosIn.getSelectedFile();

             if (fileIn != null){ 
                
                System.out.println(ejecutablePath+" -op=e -in="+fileIn);
                String resultado=ejecutar(ejecutablePath+" -op=e -in="+fileIn);
                String[] arr=resultado.split("\n");
                if (arr[0]=="true"){
                   JOptionPane.showMessageDialog(new JFrame(), "Contiene error en el bloque "+arr[1]+" en la posición "+arr[2]+".", "Comprobar error", JOptionPane.ERROR_MESSAGE);
                }else{
                   JOptionPane.showMessageDialog(new JFrame(),"No contiene error "+fileIn.getName() , "Comprobar error", JOptionPane.INFORMATION_MESSAGE);
                }
            }
    }//GEN-LAST:event_botonComprobarActionPerformed

    private void botonDañarActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonDañarActionPerformed
        // TODO add your handling code here:
        File fileIn = this.seletorArchivosIn.getSelectedFile();
           if (fileIn != null){  
                JPanel panel = new JPanel();
                JFileChooser fc = new JFileChooser();
                panel.add(fc);
                fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
                fc.setApproveButtonText("Dañar");
                panel.setVisible(true);
                int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
                if (seleccion == JFileChooser.APPROVE_OPTION){
                    File fileOut = fc.getSelectedFile();
                     if (fileOut != null){
                        System.out.println(ejecutablePath+" -op=i -in="+fileIn+" -out="+fileOut);
                         String resultado=ejecutar(ejecutablePath+" -op=i -in="+fileIn+" -out="+fileOut);
                        if (resultado.length() >1){
                            JOptionPane.showMessageDialog(new JFrame(),resultado,"Introducir error", JOptionPane.ERROR_MESSAGE);
                        }
                        seletorArchivosIn.updateUI();
                    }
                } 
              
          }
    }//GEN-LAST:event_botonDañarActionPerformed

    private void botonCorregirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonCorregirActionPerformed
        // TODO add your handling code here:
        File fileIn = this.seletorArchivosIn.getSelectedFile();
           if (fileIn != null){  
                JPanel panel = new JPanel();
                JFileChooser fc = new JFileChooser();
                panel.add(fc);
                fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
                fc.setApproveButtonText("Dañar");
                panel.setVisible(true);
                int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
                if (seleccion == JFileChooser.APPROVE_OPTION){
                    File fileOut = fc.getSelectedFile();
                     if (fileOut != null){
                        System.out.println(ejecutablePath+" -op=r -in="+fileIn+" -out="+fileOut);
                         String resultado=ejecutar(ejecutablePath+" -op=r -in="+fileIn+" -out="+fileOut);
                        if (resultado.length() >1){
                            JOptionPane.showMessageDialog(new JFrame(),resultado,"Corregir error", JOptionPane.ERROR_MESSAGE);
                        }
                        seletorArchivosIn.updateUI();
                    }
                } 
              
          }
        
    }//GEN-LAST:event_botonCorregirActionPerformed

    private void botonVerMatricesActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonVerMatricesActionPerformed
        // TODO add your handling code here:
        File fileIn = this.seletorArchivosIn.getCurrentDirectory();
        if (fileIn != null){
            System.out.println(ejecutablePath+" -op=m -out="+fileIn);
            String resultado=ejecutar(ejecutablePath+" -op=m -out="+fileIn);
            if (resultado.length() >1){
                JOptionPane.showMessageDialog(new JFrame(),resultado,"Ver Matrices", JOptionPane.ERROR_MESSAGE);
            }
            seletorArchivosIn.updateUI();
        
        }
        
    }//GEN-LAST:event_botonVerMatricesActionPerformed

    private void seletorArchivosInActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_seletorArchivosInActionPerformed
        // TODO add your handling code here:
        if ("ApproveSelection".equals(evt.getActionCommand())){
            File fileIn = this.seletorArchivosIn.getSelectedFile();
            if (fileIn != null){  
                String ext = getExtension(fileIn.getName());
                switch (ext){
                        case "ham":
                            System.out.println("Ham"); 
                        case "huf":
                            System.out.println("Huf"); 
                        case "txt":
                        case "xml":
                        case "json":
                            try{
                                Process p=Runtime.getRuntime().exec(editorPreferido+" "+fileIn.getAbsolutePath());
                                p.waitFor(); 
                            }
                            catch( Exception e){
                                System.out.println(e.toString());
                            }
                            break;
                        default:
                            System.out.println("No se puede abrir");
                            break;
                }
            }

        }
        if (evt.getActionCommand().equals("CancelSelection")){
            System.out.println("Cancel");
        }
    }//GEN-LAST:event_seletorArchivosInActionPerformed

    private void botonEditorActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonEditorActionPerformed
        // TODO add your handling code here:
       
        JPanel panel = new JPanel();
        JFileChooser fc = new JFileChooser();
        panel.add(fc);
        fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
        fc.setFileFilter(new FileFilter() {
            @Override
            public boolean accept(File f) {
                if ( getExtension(f.getName()).equals("exe")){
                    return true;
                }else{
                    if (f.isDirectory()){
                    return true;
                    }
                }
                return false;
            }

            @Override
            public String getDescription() {
                return "Seleccionar archivos .exe como editor de texto.";
            }
        });
        fc.setApproveButtonText("Editor");
        panel.setVisible(true);
        int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
        if (seleccion == JFileChooser.APPROVE_OPTION){
            File fileOut = fc.getSelectedFile();
             if (fileOut != null){  
                 editorPreferido=fileOut.getAbsolutePath();
            }
        } 
    }//GEN-LAST:event_botonEditorActionPerformed

    private void formWindowClosed(java.awt.event.WindowEvent evt) {//GEN-FIRST:event_formWindowClosed
        // TODO add your handling code here: 
    }//GEN-LAST:event_formWindowClosed

    private void formWindowClosing(java.awt.event.WindowEvent evt) {//GEN-FIRST:event_formWindowClosing
        // TODO add your handling code here: 
        System.out.println("Cierro la ventana");
        guardarPreferencias();
    }//GEN-LAST:event_formWindowClosing

    private void botonRestablecerActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonRestablecerActionPerformed
        restablercerPreferencias();
    }//GEN-LAST:event_botonRestablecerActionPerformed

    private void botonEjecutableActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonEjecutableActionPerformed
        // TODO add your handling code here:
        JPanel panel = new JPanel();
        JFileChooser fc = new JFileChooser();
        panel.add(fc);
        fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
        fc.setFileFilter(new FileFilter() {
            @Override
            public boolean accept(File f) {
                if ( getExtension(f.getName()).equals("exe")){
                    return true;
                }else{
                    if (f.isDirectory()){
                    return true;
                    }
                }
                return false;
            }
            @Override
            public String getDescription() {
                return "Seleccionar archivos .exe como editor de texto.";
            }
        });
        fc.setApproveButtonText("Ejecutable");
        panel.setVisible(true);
        int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
        if (seleccion == JFileChooser.APPROVE_OPTION){
            File fileOut = fc.getSelectedFile();
             if (fileOut != null){  
                 ejecutablePath=fileOut.getAbsolutePath();
            }
        } 
    }//GEN-LAST:event_botonEjecutableActionPerformed

    private void verConfiguracionActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_verConfiguracionActionPerformed
            
        JOptionPane.showMessageDialog(new JFrame(),"Editor: "+editorPreferido+"\nEjecutable: "+ejecutablePath,"Configuración", JOptionPane.INFORMATION_MESSAGE);
        
        
    }//GEN-LAST:event_verConfiguracionActionPerformed
   public String getExtension(String path){
        String extension = "";
        int i= path.lastIndexOf(".");
        if (i>0){
            extension = path.substring(i+1);
        }
        return extension ;
    }
    /**
     * @param args the command line arguments
     */
    public static void main(String args[]) {
        /* Set the Nimbus look and feel */
        //<editor-fold defaultstate="collapsed" desc=" Look and feel setting code (optional) ">
        /* If Nimbus (introduced in Java SE 6) is not available, stay with the default look and feel.
         * For details see http://download.oracle.com/javase/tutorial/uiswing/lookandfeel/plaf.html 
         */
        try {
            for (javax.swing.UIManager.LookAndFeelInfo info : javax.swing.UIManager.getInstalledLookAndFeels()) {
                if ("Nimbus".equals(info.getName())) {
                    javax.swing.UIManager.setLookAndFeel(info.getClassName());
                    break;
                }
            }
        } catch (ClassNotFoundException ex) {
            java.util.logging.Logger.getLogger(VentanaPrincipal.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (InstantiationException ex) {
            java.util.logging.Logger.getLogger(VentanaPrincipal.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (IllegalAccessException ex) {
            java.util.logging.Logger.getLogger(VentanaPrincipal.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (javax.swing.UnsupportedLookAndFeelException ex) {
            java.util.logging.Logger.getLogger(VentanaPrincipal.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        }
        //</editor-fold>

        /* Create and display the form */
        java.awt.EventQueue.invokeLater(new Runnable() {
            public void run() {
                new VentanaPrincipal().setVisible(true);
            }
        });
    }

    // Variables declaration - do not modify//GEN-BEGIN:variables
    private javax.swing.JMenuBar BarraMenu;
    private javax.swing.JButton botonComprimir;
    private javax.swing.JButton botonComprobar;
    private javax.swing.JButton botonCorregir;
    private javax.swing.JButton botonDañar;
    private javax.swing.JButton botonDescomprimir;
    private javax.swing.JButton botonDesproteger;
    private javax.swing.JMenuItem botonEditor;
    private javax.swing.JMenuItem botonEjecutable;
    private javax.swing.JButton botonProteger;
    private javax.swing.JMenuItem botonRestablecer;
    private javax.swing.JButton botonVerArbol;
    private javax.swing.JButton botonVerMatrices;
    private javax.swing.JMenu jMenu1;
    private javax.swing.JMenu jMenu2;
    private javax.swing.JPanel jPanel1;
    private javax.swing.JPanel jPanel2;
    private javax.swing.JMenu menuCodificacion;
    private javax.swing.JFileChooser seletorArchivosIn;
    private javax.swing.JMenuItem verConfiguracion;
    // End of variables declaration//GEN-END:variables

    private static class FileViewImpl extends FileView {

        public FileViewImpl() {
        }
    }
}
