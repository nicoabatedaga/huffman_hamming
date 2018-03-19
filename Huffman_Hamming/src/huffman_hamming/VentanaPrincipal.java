/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package huffman_hamming;

import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.Toolkit;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.ItemEvent;
import java.awt.event.ItemListener;
import java.awt.image.BufferedImage;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.URL;
import java.util.Properties;
import java.util.Vector;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.prefs.Preferences;
import javax.imageio.ImageIO;
import javax.swing.Action;
import javax.swing.ButtonGroup;
import javax.swing.DefaultComboBoxModel;
import javax.swing.ImageIcon;
import javax.swing.JComboBox;
import javax.swing.JFileChooser;
import javax.swing.JFrame;
import javax.swing.JLabel;
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
        this.setIconImage(new ImageIcon(getClass().getResource("\\images\\unsl.ico")).getImage());

        this.botonComprimir.setEnabled(true);
        this.botonComprobar.setEnabled(false);           
        this.botonCorregir.setEnabled(false);
        this.botonDañar.setEnabled(false);
        this.botonDescomprimir.setEnabled(false);
        this.botonDesproteger.setEnabled(false);
        this.botonProteger.setEnabled(true);
        this.botonComprimir.setVisible(true);
        this.botonComprobar.setVisible(false);           
        this.botonCorregir.setVisible(false);
        this.botonDañar.setVisible(false);
        this.botonDescomprimir.setVisible(false);
        this.botonDesproteger.setVisible(false);
        this.botonProteger.setVisible(true);
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
        botonProteger = new javax.swing.JButton();
        botonDesproteger = new javax.swing.JButton();
        botonComprobar = new javax.swing.JButton();
        botonDañar = new javax.swing.JButton();
        botonCorregir = new javax.swing.JButton();
        BarraMenu = new javax.swing.JMenuBar();
        botonPreferencias = new javax.swing.JMenu();
        botonEditor = new javax.swing.JMenuItem();
        menuCodificacion = new javax.swing.JMenu();
        verConfiguracion = new javax.swing.JMenuItem();
        botonRestablecer = new javax.swing.JMenuItem();
        jSeparator1 = new javax.swing.JPopupMenu.Separator();
        botonSalir = new javax.swing.JMenuItem();
        botonInformacion = new javax.swing.JMenu();
        botonAcerca = new javax.swing.JMenuItem();

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        setIconImage(Toolkit.getDefaultToolkit().getImage("\\images\\unsl.ico"));
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
                .addComponent(botonProteger, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonDesproteger, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonComprobar, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonDañar, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonCorregir, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addContainerGap(218, Short.MAX_VALUE))
        );
        jPanel2Layout.setVerticalGroup(
            jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel2Layout.createSequentialGroup()
                .addContainerGap()
                .addGroup(jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                    .addGroup(jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.BASELINE)
                        .addComponent(botonDescomprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonProteger, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonDesproteger, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonComprobar, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonDañar, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonCorregir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE))
                    .addComponent(botonComprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE))
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

        botonPreferencias.setText("Preferencias");

        botonEditor.setText("Editor de texto");
        botonEditor.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonEditorActionPerformed(evt);
            }
        });
        botonPreferencias.add(botonEditor);

        menuCodificacion.setText("Codificación");
        botonPreferencias.add(menuCodificacion);

        verConfiguracion.setText("Ver configuración");
        verConfiguracion.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                verConfiguracionActionPerformed(evt);
            }
        });
        botonPreferencias.add(verConfiguracion);

        botonRestablecer.setText("Restabler configuración");
        botonRestablecer.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonRestablecerActionPerformed(evt);
            }
        });
        botonPreferencias.add(botonRestablecer);
        botonPreferencias.add(jSeparator1);

        botonSalir.setText("Salir");
        botonSalir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonSalirActionPerformed(evt);
            }
        });
        botonPreferencias.add(botonSalir);

        BarraMenu.add(botonPreferencias);

        botonInformacion.setText("Información");
        botonInformacion.addMenuListener(new javax.swing.event.MenuListener() {
            public void menuCanceled(javax.swing.event.MenuEvent evt) {
            }
            public void menuDeselected(javax.swing.event.MenuEvent evt) {
            }
            public void menuSelected(javax.swing.event.MenuEvent evt) {
                botonInformacionMenuSelected(evt);
            }
        });
        botonInformacion.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonInformacionActionPerformed(evt);
            }
        });

        botonAcerca.setText("Acerca");
        botonAcerca.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonAcercaActionPerformed(evt);
            }
        });
        botonInformacion.add(botonAcerca);

        BarraMenu.add(botonInformacion);

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
                System.out.println(ejecutablePath+" -op=p -in="+fileIn+" -out="+convertirHam(fileOut.getAbsolutePath())+" -cod="+codificacion);
                 ejecutar(ejecutablePath+" -op=p -in="+fileIn+" -out="+convertirHam(fileOut.getAbsolutePath())+" -cod="+codificacion);
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
                    this.botonComprimir.setVisible(false);
                    this.botonComprobar.setVisible(true);
                    this.botonCorregir.setVisible(true);
                    this.botonDañar.setVisible(true);
                    this.botonDescomprimir.setVisible(false);
                    this.botonDesproteger.setVisible(true);
                    this.botonProteger.setVisible(false);
                    break;
                case "huf":
                    this.botonComprimir.setEnabled(false);
                    this.botonComprobar.setEnabled(false);
                    this.botonCorregir.setEnabled(false);
                    this.botonDañar.setEnabled(false);
                    this.botonDescomprimir.setEnabled(true);
                    this.botonDesproteger.setEnabled(false);
                    this.botonProteger.setEnabled(true);
                    this.botonComprimir.setVisible(false);
                    this.botonComprobar.setVisible(false);
                    this.botonCorregir.setVisible(false);
                    this.botonDañar.setVisible(false);
                    this.botonDescomprimir.setVisible(true);
                    this.botonDesproteger.setVisible(false);
                    this.botonProteger.setVisible(true);
                    break;
                default:
                    this.botonComprimir.setEnabled(true);
                    this.botonComprobar.setEnabled(false);           
                    this.botonCorregir.setEnabled(false);
                    this.botonDañar.setEnabled(false);
                    this.botonDescomprimir.setEnabled(false);
                    this.botonDesproteger.setEnabled(false);
                    this.botonProteger.setEnabled(true);
                    this.botonComprimir.setVisible(true);
                    this.botonComprobar.setVisible(false);           
                    this.botonCorregir.setVisible(false);
                    this.botonDañar.setVisible(false);
                    this.botonDescomprimir.setVisible(false);
                    this.botonDesproteger.setVisible(false);
                    this.botonProteger.setVisible(true);
                        
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
                if (arr[0].equals("true")){
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
                        System.out.println(ejecutablePath+" -op=i -in="+fileIn+" -out="+convertirHam(fileOut.getAbsolutePath()));
                         String resultado=ejecutar(ejecutablePath+" -op=i -in="+fileIn+" -out="+convertirHam(fileOut.getAbsolutePath()));
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
                        System.out.println(ejecutablePath+" -op=r -in="+fileIn+" -out="+convertirHam(fileOut.getAbsolutePath()));
                         String resultado=ejecutar(ejecutablePath+" -op=r -in="+fileIn+" -out="+convertirHam(fileOut.getAbsolutePath()));
                        if (resultado.length() >1){
                            JOptionPane.showMessageDialog(new JFrame(),resultado,"Corregir error", JOptionPane.ERROR_MESSAGE);
                        }
                        seletorArchivosIn.updateUI();
                    }
                } 
              
          }
        
    }//GEN-LAST:event_botonCorregirActionPerformed
public void mostrarImagen(String path,String tittle,int x,int y){
    try {
    final JFrame showPictureFrame = new JFrame(tittle);
    JLabel pictureLabel = new JLabel();
    URL url = new File(path).toURI().toURL();
    BufferedImage img = ImageIO.read(url);
    pictureLabel.setIcon(new ImageIcon(img));
    // add the label to the frame
    showPictureFrame.add(pictureLabel);
    showPictureFrame.setLocation(x,y);
    // pack everything (does many stuff. e.g. resizes the frame to fit the image)
    showPictureFrame.pack();

    //this is how you should open a new Frame or Dialog, but only using showPictureFrame.setVisible(true); would also work.
    java.awt.EventQueue.invokeLater(new Runnable() {

      public void run() {
        showPictureFrame.setVisible(true);
      }
    });

  } catch (IOException ex) {
    System.err.println("Some IOException accured (did you set the right path?): ");
    System.err.println(ex.getMessage());
  }
}
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

    private void verConfiguracionActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_verConfiguracionActionPerformed
        JOptionPane.showMessageDialog(new JFrame(),"Editor: "+editorPreferido+"\nEjecutable: "+ejecutablePath,"Configuración", JOptionPane.INFORMATION_MESSAGE);
        
    }//GEN-LAST:event_verConfiguracionActionPerformed

    private void botonSalirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonSalirActionPerformed
        // TODO add your handling code here:
        guardarPreferencias();
        System.exit(0);
    }//GEN-LAST:event_botonSalirActionPerformed

    private void botonInformacionActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonInformacionActionPerformed
        // TODO add your handling code here: 
    }//GEN-LAST:event_botonInformacionActionPerformed

    private void botonInformacionMenuSelected(javax.swing.event.MenuEvent evt) {//GEN-FIRST:event_botonInformacionMenuSelected


        // TODO add your handling code here:
    }//GEN-LAST:event_botonInformacionMenuSelected

    private void botonAcercaActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonAcercaActionPerformed
        // TODO add your handling code here:
        System.out.println("Acerca");
         final ImageIcon icon = new ImageIcon("C:\\Users\\Joaquin\\go\\src\\huffman_hamming\\Huffman_Hamming\\src\\huffman_hamming\\images\\unsl.ico");
          
        JOptionPane.showMessageDialog(null,"Aplicación desarrollada en el ambito\n de la materia Teoría de la información\n para la carrera Ing. en informatica. ","Acerca de Huffman Hamming", JOptionPane.DEFAULT_OPTION,icon);

    }//GEN-LAST:event_botonAcercaActionPerformed
    public String convertirHam(String path){
        if (getExtension(path).equals("ham")){
            return path;
        }else{
            return path+".ham";
        }
    }
    
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
                System.out.println("Nombre:"+info.getName());
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
    private javax.swing.JMenuItem botonAcerca;
    private javax.swing.JButton botonComprimir;
    private javax.swing.JButton botonComprobar;
    private javax.swing.JButton botonCorregir;
    private javax.swing.JButton botonDañar;
    private javax.swing.JButton botonDescomprimir;
    private javax.swing.JButton botonDesproteger;
    private javax.swing.JMenuItem botonEditor;
    private javax.swing.JMenu botonInformacion;
    private javax.swing.JMenu botonPreferencias;
    private javax.swing.JButton botonProteger;
    private javax.swing.JMenuItem botonRestablecer;
    private javax.swing.JMenuItem botonSalir;
    private javax.swing.JPanel jPanel1;
    private javax.swing.JPanel jPanel2;
    private javax.swing.JPopupMenu.Separator jSeparator1;
    private javax.swing.JMenu menuCodificacion;
    private javax.swing.JFileChooser seletorArchivosIn;
    private javax.swing.JMenuItem verConfiguracion;
    // End of variables declaration//GEN-END:variables

    private static class FileViewImpl extends FileView {

        public FileViewImpl() {
        }
    }
}
