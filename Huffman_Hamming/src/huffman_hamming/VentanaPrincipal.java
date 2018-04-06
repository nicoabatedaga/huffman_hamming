/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package huffman_hamming;
  
import java.awt.Desktop;
import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.Image;
import java.awt.Toolkit;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.ItemEvent;
import java.awt.event.ItemListener;
import java.awt.image.BufferedImage;
import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.URL;  
import java.nio.file.Files;
import java.nio.file.attribute.BasicFileAttributes;
import java.util.LinkedList;
import java.util.List;
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
import javax.swing.JButton;
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
    File tempFile;
    private String ejecutablePath=getClass().getResource("huffman_hamming.exe").getPath(); 
   // private String ejecutablePath="C:\\Users\\Joaquin\\go\\src\\huffman_hamming\\huffman_hamming.exe"; 
    public String convertirHam(String path){
        if (getExtension(path).equals("ham")){
            return path;
        }else{
            return path+".ham";
        }
    }
     public String convertirHuff(String path){
        if (getExtension(path).equals("huff")){
            return path;
        }else{
            return path+".huff";
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
    public void cargarPreferencias(){
        Properties prop = new Properties();
        InputStream input = null;
        try {
            input = new FileInputStream("./preferencias");
            prop.load(input);
            editorPreferido=(prop.getProperty("editor")); 
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
    public void restablercerPreferencias(){
        try{    
            String EDITOR = "notepad"; 
             //create a properties file
             Properties props = new Properties();
             props.setProperty("editor", EDITOR);
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
        //create a properties file
        Properties props = new Properties();
        props.setProperty("editor", EDITOR); 
        File f = new File("./preferencias");
        OutputStream out = new FileOutputStream( f );
        //If you wish to make some comments 
        props.store(out, "Preferencias");
   }
    catch (Exception e ) {
        e.printStackTrace();
    }
   }
   private void streamToFile(InputStream ip,String path) throws IOException{
       File file= new File(path);
       OutputStream os = null;  
        try {   
            os = new FileOutputStream(file);
            byte[] buffer = new byte[1024];
            int length;
            while ((length = ip.read(buffer)) > 0) {
                os.write(buffer, 0, length);
            }
        } catch (FileNotFoundException ex) {
            Logger.getLogger(VentanaPrincipal.class.getName()).log(Level.SEVERE, null, ex);
        } finally {
            ip.close();
            os.close();
        }
 
   }
   
    public VentanaPrincipal() {
        try {
            this.tempFile = File.createTempFile("huffman_hamming", ".exe");
            streamToFile(this.getClass().getResourceAsStream("huffman_hamming.exe"),this.tempFile.getPath());
            this.tempFile.deleteOnExit();
        } catch (IOException ex) {
            Logger.getLogger(VentanaPrincipal.class.getName()).log(Level.SEVERE, null, ex);
        }
        this.ejecutablePath = this.tempFile.getPath();
        initComponents();  
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
       this.setIconImage(Toolkit.getDefaultToolkit().getImage(getClass().getResource("unsl.png")));

    }

    /**
     * This method is called from within the constructor to initialize the form.
     * WARNING: Do NOT modify this code. The content of this method is always
     * regenerated by the Form Editor.
     */
    @SuppressWarnings("unchecked")
    // <editor-fold defaultstate="collapsed" desc="Generated Code">//GEN-BEGIN:initComponents
    private void initComponents() {

        Otro = new javax.swing.JFrame();
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
        botonDiferencia = new javax.swing.JButton();
        filler2 = new javax.swing.Box.Filler(new java.awt.Dimension(0, 0), new java.awt.Dimension(0, 0), new java.awt.Dimension(0, 0));
        filler1 = new javax.swing.Box.Filler(new java.awt.Dimension(0, 0), new java.awt.Dimension(0, 0), new java.awt.Dimension(0, 32767));
        BarraMenu = new javax.swing.JMenuBar();
        botonPreferencias = new javax.swing.JMenu();
        botonEditor = new javax.swing.JMenuItem();
        menuCodificacion = new javax.swing.JMenu();
        jSeparator1 = new javax.swing.JPopupMenu.Separator();
        botonAcerca = new javax.swing.JMenuItem();
        jSeparator2 = new javax.swing.JPopupMenu.Separator();
        botonSalir = new javax.swing.JMenuItem();

        this.setIconImage(Toolkit.getDefaultToolkit().getImage("unsl.png"));

        javax.swing.GroupLayout OtroLayout = new javax.swing.GroupLayout(Otro.getContentPane());
        Otro.getContentPane().setLayout(OtroLayout);
        OtroLayout.setHorizontalGroup(
            OtroLayout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGap(0, 400, Short.MAX_VALUE)
        );
        OtroLayout.setVerticalGroup(
            OtroLayout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGap(0, 300, Short.MAX_VALUE)
        );

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        setTitle("Huffman Hamming");
        setResizable(false);
        addWindowListener(new java.awt.event.WindowAdapter() {
            public void windowClosed(java.awt.event.WindowEvent evt) {
                formWindowClosed(evt);
            }
            public void windowClosing(java.awt.event.WindowEvent evt) {
                formWindowClosing(evt);
            }
        });

        jPanel1.setPreferredSize(new java.awt.Dimension(800, 450));

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

        jPanel2.setPreferredSize(new java.awt.Dimension(915, 51));

        botonComprimir.setText("Comprimir");
        botonComprimir.setMaximumSize(new java.awt.Dimension(0, 0));
        botonComprimir.setMinimumSize(new java.awt.Dimension(0, 0));
        botonComprimir.setPreferredSize(new java.awt.Dimension(100, 35));
        botonComprimir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonComprimirActionPerformed(evt);
            }
        });

        botonDescomprimir.setText("Descomprimir");
        botonDescomprimir.setMaximumSize(new java.awt.Dimension(0, 0));
        botonDescomprimir.setMinimumSize(new java.awt.Dimension(0, 0));
        botonDescomprimir.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDescomprimir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDescomprimirActionPerformed(evt);
            }
        });

        botonProteger.setText("Proteger");
        botonProteger.setMaximumSize(new java.awt.Dimension(0, 0));
        botonProteger.setMinimumSize(new java.awt.Dimension(0, 0));
        botonProteger.setPreferredSize(new java.awt.Dimension(100, 35));
        botonProteger.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonProtegerActionPerformed(evt);
            }
        });

        botonDesproteger.setText("Desproteger");
        botonDesproteger.setMaximumSize(new java.awt.Dimension(0, 0));
        botonDesproteger.setMinimumSize(new java.awt.Dimension(0, 0));
        botonDesproteger.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDesproteger.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDesprotegerActionPerformed(evt);
            }
        });

        botonComprobar.setText("Comprobar");
        botonComprobar.setMaximumSize(new java.awt.Dimension(0, 0));
        botonComprobar.setMinimumSize(new java.awt.Dimension(0, 0));
        botonComprobar.setPreferredSize(new java.awt.Dimension(100, 35));
        botonComprobar.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonComprobarActionPerformed(evt);
            }
        });

        botonDañar.setText("Dañar");
        botonDañar.setMaximumSize(new java.awt.Dimension(0, 0));
        botonDañar.setMinimumSize(new java.awt.Dimension(0, 0));
        botonDañar.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDañar.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDañarActionPerformed(evt);
            }
        });

        botonCorregir.setText("Corregir");
        botonCorregir.setMaximumSize(new java.awt.Dimension(0, 0));
        botonCorregir.setMinimumSize(new java.awt.Dimension(0, 0));
        botonCorregir.setPreferredSize(new java.awt.Dimension(100, 35));
        botonCorregir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonCorregirActionPerformed(evt);
            }
        });

        botonDiferencia.setText("Diferencia");
        botonDiferencia.setMaximumSize(new java.awt.Dimension(0, 0));
        botonDiferencia.setMinimumSize(new java.awt.Dimension(0, 0));
        botonDiferencia.setPreferredSize(new java.awt.Dimension(100, 35));
        botonDiferencia.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonDiferenciaActionPerformed(evt);
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
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(botonDiferencia, javax.swing.GroupLayout.PREFERRED_SIZE, 98, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(114, 114, 114)
                .addComponent(filler2, javax.swing.GroupLayout.PREFERRED_SIZE, 41, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );
        jPanel2Layout.setVerticalGroup(
            jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel2Layout.createSequentialGroup()
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                .addGroup(jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING, false)
                    .addGroup(jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.BASELINE)
                        .addComponent(botonDescomprimir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonProteger, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonDesproteger, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonComprobar, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonDañar, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonCorregir, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addComponent(botonDiferencia, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE))
                    .addComponent(botonComprimir, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                    .addComponent(filler2, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)))
        );

        javax.swing.GroupLayout jPanel1Layout = new javax.swing.GroupLayout(jPanel1);
        jPanel1.setLayout(jPanel1Layout);
        jPanel1Layout.setHorizontalGroup(
            jPanel1Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(filler1, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
            .addComponent(seletorArchivosIn, javax.swing.GroupLayout.PREFERRED_SIZE, 800, javax.swing.GroupLayout.PREFERRED_SIZE)
            .addComponent(jPanel2, javax.swing.GroupLayout.PREFERRED_SIZE, 948, javax.swing.GroupLayout.PREFERRED_SIZE)
        );
        jPanel1Layout.setVerticalGroup(
            jPanel1Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel1Layout.createSequentialGroup()
                .addComponent(jPanel2, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(seletorArchivosIn, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(filler1, javax.swing.GroupLayout.PREFERRED_SIZE, 18, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addContainerGap())
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
        botonPreferencias.add(jSeparator1);

        botonAcerca.setText("Acerca");
        botonAcerca.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonAcercaActionPerformed(evt);
            }
        });
        botonPreferencias.add(botonAcerca);
        botonPreferencias.add(jSeparator2);

        botonSalir.setText("Salir");
        botonSalir.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                botonSalirActionPerformed(evt);
            }
        });
        botonPreferencias.add(botonSalir);

        BarraMenu.add(botonPreferencias);

        setJMenuBar(BarraMenu);

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(getContentPane());
        getContentPane().setLayout(layout);
        layout.setHorizontalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jPanel1, javax.swing.GroupLayout.PREFERRED_SIZE, 806, javax.swing.GroupLayout.PREFERRED_SIZE)
        );
        layout.setVerticalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jPanel1, 466, 466, javax.swing.GroupLayout.PREFERRED_SIZE)
        );

        getAccessibleContext().setAccessibleDescription("");
        getAccessibleContext().setAccessibleParent(Otro);

        pack();
        setLocationRelativeTo(null);
    }// </editor-fold>//GEN-END:initComponents

    private void set(JButton j,boolean b){
        j.setEnabled(b);
        j.setVisible(b);
    }
    
    private void seletorArchivosInPropertyChange(java.beans.PropertyChangeEvent evt) {//GEN-FIRST:event_seletorArchivosInPropertyChange
    if(JFileChooser.SELECTED_FILE_CHANGED_PROPERTY.equals(evt.getPropertyName())){
        File file =this.seletorArchivosIn.getSelectedFile();
        if (file != null){
            String extension =getExtension(file.toString());
            this.botonDiferencia.setVisible(true); 
            switch (extension){
                case "ham":
                    set(this.botonComprimir,true);
                    set(this.botonComprobar,true);
                    set(this.botonCorregir,true);
                    set(this.botonDañar,true);
                    set(this.botonDescomprimir,false);
                    set(this.botonDesproteger,true);
                    set(this.botonProteger,false);
                    set(this.botonDiferencia,false);
                    break;
                case "huff":
                    set(this.botonComprimir,false);
                    set(this.botonComprobar,false);
                    set(this.botonCorregir,false);
                    set(this.botonDañar,false);
                    set(this.botonDescomprimir,true);
                    set(this.botonDesproteger,false);
                    set(this.botonProteger,true);
                    set(this.botonDiferencia,false); 
                    break;
                case "txt":
                    set(this.botonComprimir,true);
                    set(this.botonComprobar,false);
                    set(this.botonCorregir,false);
                    set(this.botonDañar,false);
                    set(this.botonDescomprimir,false);
                    set(this.botonDesproteger,false);
                    set(this.botonProteger,true);
                    set(this.botonDiferencia,true);  
                    break;
                default:
                    set(this.botonComprimir,true);
                    set(this.botonComprobar,false);
                    set(this.botonCorregir,false);
                    set(this.botonDañar,false);
                    set(this.botonDescomprimir,false);
                    set(this.botonDesproteger,false);
                    set(this.botonProteger,true);
                    set(this.botonDiferencia,false);  
                    break; 
            } 
        }else{ 
            set(this.botonComprimir,false);
            set(this.botonComprobar,false);
            set(this.botonCorregir,false);
            set(this.botonDañar,false);
            set(this.botonDescomprimir,false);
            set(this.botonDesproteger,false);
            set(this.botonProteger,false);
            set(this.botonDiferencia,false);   
        }
    }       
        
// TODO add your handling code here:
    }//GEN-LAST:event_seletorArchivosInPropertyChange

    public String getStringErrores(String[] arr ){
        String resultado="";
        for (int i =0;i<arr.length;i+=2){
            resultado+="("+arr[i]+" , "+arr[i+1]+"), ";
            if ((i+1)%5 == 0){
                resultado+="\n";
            }
         }
        return resultado;
    }
    
    private void seletorArchivosInActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_seletorArchivosInActionPerformed
        // TODO add your handling code here:
        if ("ApproveSelection".equals(evt.getActionCommand())){
            File fileIn = this.seletorArchivosIn.getSelectedFile();
            if (fileIn != null){  
                String ext = getExtension(fileIn.getName());
                switch (ext){
                        case "ham": 
                        case "huf": 
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
                            break;
                }
            }

        }
        if (evt.getActionCommand().equals("CancelSelection")){
            System.exit(0); 
        }
    }//GEN-LAST:event_seletorArchivosInActionPerformed

    private void formWindowClosed(java.awt.event.WindowEvent evt) {//GEN-FIRST:event_formWindowClosed
        // TODO add your handling code here: 
    }//GEN-LAST:event_formWindowClosed

    private void formWindowClosing(java.awt.event.WindowEvent evt) {//GEN-FIRST:event_formWindowClosing
        // TODO add your handling code here:  
        guardarPreferencias();
    }//GEN-LAST:event_formWindowClosing

    private void botonAcercaActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonAcercaActionPerformed
        // TODO add your handling code here: 
        JOptionPane.showMessageDialog(null,"<html><b>Huffman Hamming</b>, aplicación desarrollada en el ambito de la <br>matería Teoría de la Información para la carrera Ing. en Informatica.<br><br><b>Desarrollada por:</b><br><i>Abatedaga Biole, Nicolas y Loyola, Franco Joaquín</i></body></html> ","Huffman Hamming", JOptionPane.DEFAULT_OPTION);

    }//GEN-LAST:event_botonAcercaActionPerformed

    private void botonSalirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonSalirActionPerformed
        // TODO add your handling code here:
        guardarPreferencias();
        System.exit(0);
    }//GEN-LAST:event_botonSalirActionPerformed

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

    private void botonDiferenciaActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonDiferenciaActionPerformed
        File fileIn = this.seletorArchivosIn.getSelectedFile();
        if (fileIn != null){
            JPanel panel = new JPanel();
            JFileChooser fc = new JFileChooser();
            panel.add(fc);
            fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
            fc.setApproveButtonText("Comparar");
            panel.setVisible(true);
            int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
            if (seleccion == JFileChooser.APPROVE_OPTION){
                File fileOut = fc.getSelectedFile();
                String nombreCarpeta="/diff-"+fileIn.getName();
                if (nombreCarpeta.contains(".")){
                    nombreCarpeta = nombreCarpeta.replace('.','-');
                }
                nombreCarpeta+="/";
                String dir=fileOut.getParent()+nombreCarpeta;
                File fdir = new File(dir);
                if (!fdir.exists()){
                    fdir.mkdir();
                } 
                String dirJs=dir+"/js";
                File jsdir = new File(dirJs);
                if (!jsdir.exists()){
                    jsdir.mkdir();
                }
             
                stringToFile("function diff_match_patch(){this.Diff_Timeout=1;this.Diff_EditCost=4;this.Match_Threshold=.5;this.Match_Distance=1E3;this.Patch_DeleteThreshold=.5;this.Patch_Margin=4;this.Match_MaxBits=32}var DIFF_DELETE=-1,DIFF_INSERT=1,DIFF_EQUAL=0;\n" +
"diff_match_patch.prototype.diff_main=function(a,b,c,d){\"undefined\"==typeof d&&(d=0>=this.Diff_Timeout?Number.MAX_VALUE:(new Date).getTime()+1E3*this.Diff_Timeout);if(null==a||null==b)throw Error(\"Null input. (diff_main)\");if(a==b)return a?[[DIFF_EQUAL,a]]:[];\"undefined\"==typeof c&&(c=!0);var e=c,f=this.diff_commonPrefix(a,b);c=a.substring(0,f);a=a.substring(f);b=b.substring(f);f=this.diff_commonSuffix(a,b);var g=a.substring(a.length-f);a=a.substring(0,a.length-f);b=b.substring(0,b.length-f);a=this.diff_compute_(a,\n" +
"b,e,d);c&&a.unshift([DIFF_EQUAL,c]);g&&a.push([DIFF_EQUAL,g]);this.diff_cleanupMerge(a);return a};\n" +
"diff_match_patch.prototype.diff_compute_=function(a,b,c,d){if(!a)return[[DIFF_INSERT,b]];if(!b)return[[DIFF_DELETE,a]];var e=a.length>b.length?a:b,f=a.length>b.length?b:a,g=e.indexOf(f);return-1!=g?(c=[[DIFF_INSERT,e.substring(0,g)],[DIFF_EQUAL,f],[DIFF_INSERT,e.substring(g+f.length)]],a.length>b.length&&(c[0][0]=c[2][0]=DIFF_DELETE),c):1==f.length?[[DIFF_DELETE,a],[DIFF_INSERT,b]]:(e=this.diff_halfMatch_(a,b))?(b=e[1],f=e[3],a=e[4],e=this.diff_main(e[0],e[2],c,d),c=this.diff_main(b,f,c,d),e.concat([[DIFF_EQUAL,\n" +
"a]],c)):c&&100<a.length&&100<b.length?this.diff_lineMode_(a,b,d):this.diff_bisect_(a,b,d)};\n" +
"diff_match_patch.prototype.diff_lineMode_=function(a,b,c){var d=this.diff_linesToChars_(a,b);a=d.chars1;b=d.chars2;d=d.lineArray;a=this.diff_main(a,b,!1,c);this.diff_charsToLines_(a,d);this.diff_cleanupSemantic(a);a.push([DIFF_EQUAL,\"\"]);for(var e=d=b=0,f=\"\",g=\"\";b<a.length;){switch(a[b][0]){case DIFF_INSERT:e++;g+=a[b][1];break;case DIFF_DELETE:d++;f+=a[b][1];break;case DIFF_EQUAL:if(1<=d&&1<=e){a.splice(b-d-e,d+e);b=b-d-e;d=this.diff_main(f,g,!1,c);for(e=d.length-1;0<=e;e--)a.splice(b,0,d[e]);b+=\n" +
"d.length}d=e=0;g=f=\"\"}b++}a.pop();return a};\n" +
"diff_match_patch.prototype.diff_bisect_=function(a,b,c){for(var d=a.length,e=b.length,f=Math.ceil((d+e)/2),g=2*f,h=Array(g),l=Array(g),k=0;k<g;k++)h[k]=-1,l[k]=-1;h[f+1]=0;l[f+1]=0;k=d-e;for(var m=0!=k%2,p=0,x=0,w=0,q=0,t=0;t<f&&!((new Date).getTime()>c);t++){for(var v=-t+p;v<=t-x;v+=2){var n=f+v;var r=v==-t||v!=t&&h[n-1]<h[n+1]?h[n+1]:h[n-1]+1;for(var y=r-v;r<d&&y<e&&a.charAt(r)==b.charAt(y);)r++,y++;h[n]=r;if(r>d)x+=2;else if(y>e)p+=2;else if(m&&(n=f+k-v,0<=n&&n<g&&-1!=l[n])){var u=d-l[n];if(r>=\n" +
"u)return this.diff_bisectSplit_(a,b,r,y,c)}}for(v=-t+w;v<=t-q;v+=2){n=f+v;u=v==-t||v!=t&&l[n-1]<l[n+1]?l[n+1]:l[n-1]+1;for(r=u-v;u<d&&r<e&&a.charAt(d-u-1)==b.charAt(e-r-1);)u++,r++;l[n]=u;if(u>d)q+=2;else if(r>e)w+=2;else if(!m&&(n=f+k-v,0<=n&&n<g&&-1!=h[n]&&(r=h[n],y=f+r-n,u=d-u,r>=u)))return this.diff_bisectSplit_(a,b,r,y,c)}}return[[DIFF_DELETE,a],[DIFF_INSERT,b]]};\n" +
"diff_match_patch.prototype.diff_bisectSplit_=function(a,b,c,d,e){var f=a.substring(0,c),g=b.substring(0,d);a=a.substring(c);b=b.substring(d);f=this.diff_main(f,g,!1,e);e=this.diff_main(a,b,!1,e);return f.concat(e)};\n" +
"diff_match_patch.prototype.diff_linesToChars_=function(a,b){function c(a){for(var b=\"\",c=0,f=-1,g=d.length;f<a.length-1;){f=a.indexOf(\"\\n\",c);-1==f&&(f=a.length-1);var h=a.substring(c,f+1);c=f+1;(e.hasOwnProperty?e.hasOwnProperty(h):void 0!==e[h])?b+=String.fromCharCode(e[h]):(b+=String.fromCharCode(g),e[h]=g,d[g++]=h)}return b}var d=[],e={};d[0]=\"\";var f=c(a),g=c(b);return{chars1:f,chars2:g,lineArray:d}};\n" +
"diff_match_patch.prototype.diff_charsToLines_=function(a,b){for(var c=0;c<a.length;c++){for(var d=a[c][1],e=[],f=0;f<d.length;f++)e[f]=b[d.charCodeAt(f)];a[c][1]=e.join(\"\")}};diff_match_patch.prototype.diff_commonPrefix=function(a,b){if(!a||!b||a.charAt(0)!=b.charAt(0))return 0;for(var c=0,d=Math.min(a.length,b.length),e=d,f=0;c<e;)a.substring(f,e)==b.substring(f,e)?f=c=e:d=e,e=Math.floor((d-c)/2+c);return e};\n" +
"diff_match_patch.prototype.diff_commonSuffix=function(a,b){if(!a||!b||a.charAt(a.length-1)!=b.charAt(b.length-1))return 0;for(var c=0,d=Math.min(a.length,b.length),e=d,f=0;c<e;)a.substring(a.length-e,a.length-f)==b.substring(b.length-e,b.length-f)?f=c=e:d=e,e=Math.floor((d-c)/2+c);return e};\n" +
"diff_match_patch.prototype.diff_commonOverlap_=function(a,b){var c=a.length,d=b.length;if(0==c||0==d)return 0;c>d?a=a.substring(c-d):c<d&&(b=b.substring(0,c));c=Math.min(c,d);if(a==b)return c;d=0;for(var e=1;;){var f=a.substring(c-e);f=b.indexOf(f);if(-1==f)return d;e+=f;if(0==f||a.substring(c-e)==b.substring(0,e))d=e,e++}};\n" +
"diff_match_patch.prototype.diff_halfMatch_=function(a,b){function c(a,b,c){for(var d=a.substring(c,c+Math.floor(a.length/4)),e=-1,g=\"\",h,k,l,m;-1!=(e=b.indexOf(d,e+1));){var p=f.diff_commonPrefix(a.substring(c),b.substring(e)),u=f.diff_commonSuffix(a.substring(0,c),b.substring(0,e));g.length<u+p&&(g=b.substring(e-u,e)+b.substring(e,e+p),h=a.substring(0,c-u),k=a.substring(c+p),l=b.substring(0,e-u),m=b.substring(e+p))}return 2*g.length>=a.length?[h,k,l,m,g]:null}if(0>=this.Diff_Timeout)return null;\n" +
"var d=a.length>b.length?a:b,e=a.length>b.length?b:a;if(4>d.length||2*e.length<d.length)return null;var f=this,g=c(d,e,Math.ceil(d.length/4));d=c(d,e,Math.ceil(d.length/2));if(g||d)g=d?g?g[4].length>d[4].length?g:d:d:g;else return null;if(a.length>b.length){d=g[0];e=g[1];var h=g[2];var l=g[3]}else h=g[0],l=g[1],d=g[2],e=g[3];return[d,e,h,l,g[4]]};\n" +
"diff_match_patch.prototype.diff_cleanupSemantic=function(a){for(var b=!1,c=[],d=0,e=null,f=0,g=0,h=0,l=0,k=0;f<a.length;)a[f][0]==DIFF_EQUAL?(c[d++]=f,g=l,h=k,k=l=0,e=a[f][1]):(a[f][0]==DIFF_INSERT?l+=a[f][1].length:k+=a[f][1].length,e&&e.length<=Math.max(g,h)&&e.length<=Math.max(l,k)&&(a.splice(c[d-1],0,[DIFF_DELETE,e]),a[c[d-1]+1][0]=DIFF_INSERT,d--,d--,f=0<d?c[d-1]:-1,k=l=h=g=0,e=null,b=!0)),f++;b&&this.diff_cleanupMerge(a);this.diff_cleanupSemanticLossless(a);for(f=1;f<a.length;){if(a[f-1][0]==\n" +
"DIFF_DELETE&&a[f][0]==DIFF_INSERT){b=a[f-1][1];c=a[f][1];d=this.diff_commonOverlap_(b,c);e=this.diff_commonOverlap_(c,b);if(d>=e){if(d>=b.length/2||d>=c.length/2)a.splice(f,0,[DIFF_EQUAL,c.substring(0,d)]),a[f-1][1]=b.substring(0,b.length-d),a[f+1][1]=c.substring(d),f++}else if(e>=b.length/2||e>=c.length/2)a.splice(f,0,[DIFF_EQUAL,b.substring(0,e)]),a[f-1][0]=DIFF_INSERT,a[f-1][1]=c.substring(0,c.length-e),a[f+1][0]=DIFF_DELETE,a[f+1][1]=b.substring(e),f++;f++}f++}};\n" +
"diff_match_patch.prototype.diff_cleanupSemanticLossless=function(a){function b(a,b){if(!a||!b)return 6;var c=a.charAt(a.length-1),d=b.charAt(0),e=c.match(diff_match_patch.nonAlphaNumericRegex_),f=d.match(diff_match_patch.nonAlphaNumericRegex_),g=e&&c.match(diff_match_patch.whitespaceRegex_),h=f&&d.match(diff_match_patch.whitespaceRegex_);c=g&&c.match(diff_match_patch.linebreakRegex_);d=h&&d.match(diff_match_patch.linebreakRegex_);var k=c&&a.match(diff_match_patch.blanklineEndRegex_),l=d&&b.match(diff_match_patch.blanklineStartRegex_);\n" +
"return k||l?5:c||d?4:e&&!g&&h?3:g||h?2:e||f?1:0}for(var c=1;c<a.length-1;){if(a[c-1][0]==DIFF_EQUAL&&a[c+1][0]==DIFF_EQUAL){var d=a[c-1][1],e=a[c][1],f=a[c+1][1],g=this.diff_commonSuffix(d,e);if(g){var h=e.substring(e.length-g);d=d.substring(0,d.length-g);e=h+e.substring(0,e.length-g);f=h+f}g=d;h=e;for(var l=f,k=b(d,e)+b(e,f);e.charAt(0)===f.charAt(0);){d+=e.charAt(0);e=e.substring(1)+f.charAt(0);f=f.substring(1);var m=b(d,e)+b(e,f);m>=k&&(k=m,g=d,h=e,l=f)}a[c-1][1]!=g&&(g?a[c-1][1]=g:(a.splice(c-\n" +
"1,1),c--),a[c][1]=h,l?a[c+1][1]=l:(a.splice(c+1,1),c--))}c++}};diff_match_patch.nonAlphaNumericRegex_=/[^a-zA-Z0-9]/;diff_match_patch.whitespaceRegex_=/\\s/;diff_match_patch.linebreakRegex_=/[\\r\\n]/;diff_match_patch.blanklineEndRegex_=/\\n\\r?\\n$/;diff_match_patch.blanklineStartRegex_=/^\\r?\\n\\r?\\n/;\n" +
"diff_match_patch.prototype.diff_cleanupEfficiency=function(a){for(var b=!1,c=[],d=0,e=null,f=0,g=!1,h=!1,l=!1,k=!1;f<a.length;)a[f][0]==DIFF_EQUAL?(a[f][1].length<this.Diff_EditCost&&(l||k)?(c[d++]=f,g=l,h=k,e=a[f][1]):(d=0,e=null),l=k=!1):(a[f][0]==DIFF_DELETE?k=!0:l=!0,e&&(g&&h&&l&&k||e.length<this.Diff_EditCost/2&&3==g+h+l+k)&&(a.splice(c[d-1],0,[DIFF_DELETE,e]),a[c[d-1]+1][0]=DIFF_INSERT,d--,e=null,g&&h?(l=k=!0,d=0):(d--,f=0<d?c[d-1]:-1,l=k=!1),b=!0)),f++;b&&this.diff_cleanupMerge(a)};\n" +
"diff_match_patch.prototype.diff_cleanupMerge=function(a){a.push([DIFF_EQUAL,\"\"]);for(var b=0,c=0,d=0,e=\"\",f=\"\",g;b<a.length;)switch(a[b][0]){case DIFF_INSERT:d++;f+=a[b][1];b++;break;case DIFF_DELETE:c++;e+=a[b][1];b++;break;case DIFF_EQUAL:1<c+d?(0!==c&&0!==d&&(g=this.diff_commonPrefix(f,e),0!==g&&(0<b-c-d&&a[b-c-d-1][0]==DIFF_EQUAL?a[b-c-d-1][1]+=f.substring(0,g):(a.splice(0,0,[DIFF_EQUAL,f.substring(0,g)]),b++),f=f.substring(g),e=e.substring(g)),g=this.diff_commonSuffix(f,e),0!==g&&(a[b][1]=f.substring(f.length-\n" +
"g)+a[b][1],f=f.substring(0,f.length-g),e=e.substring(0,e.length-g))),0===c?a.splice(b-d,c+d,[DIFF_INSERT,f]):0===d?a.splice(b-c,c+d,[DIFF_DELETE,e]):a.splice(b-c-d,c+d,[DIFF_DELETE,e],[DIFF_INSERT,f]),b=b-c-d+(c?1:0)+(d?1:0)+1):0!==b&&a[b-1][0]==DIFF_EQUAL?(a[b-1][1]+=a[b][1],a.splice(b,1)):b++,c=d=0,f=e=\"\"}\"\"===a[a.length-1][1]&&a.pop();c=!1;for(b=1;b<a.length-1;)a[b-1][0]==DIFF_EQUAL&&a[b+1][0]==DIFF_EQUAL&&(a[b][1].substring(a[b][1].length-a[b-1][1].length)==a[b-1][1]?(a[b][1]=a[b-1][1]+a[b][1].substring(0,\n" +
"a[b][1].length-a[b-1][1].length),a[b+1][1]=a[b-1][1]+a[b+1][1],a.splice(b-1,1),c=!0):a[b][1].substring(0,a[b+1][1].length)==a[b+1][1]&&(a[b-1][1]+=a[b+1][1],a[b][1]=a[b][1].substring(a[b+1][1].length)+a[b+1][1],a.splice(b+1,1),c=!0)),b++;c&&this.diff_cleanupMerge(a)};\n" +
"diff_match_patch.prototype.diff_xIndex=function(a,b){var c=0,d=0,e=0,f=0,g;for(g=0;g<a.length;g++){a[g][0]!==DIFF_INSERT&&(c+=a[g][1].length);a[g][0]!==DIFF_DELETE&&(d+=a[g][1].length);if(c>b)break;e=c;f=d}return a.length!=g&&a[g][0]===DIFF_DELETE?f:f+(b-e)};\n" +
"diff_match_patch.prototype.diff_prettyHtml=function(a){for(var b=[],c=/&/g,d=/</g,e=/>/g,f=/\\n/g,g=0;g<a.length;g++){var h=a[g][0],l=a[g][1].replace(c,\"&amp;\").replace(d,\"&lt;\").replace(e,\"&gt;\").replace(f,\"&para;<br>\");switch(h){case DIFF_INSERT:b[g]='<ins style=\"background:#e6ffe6;\">'+l+\"</ins>\";break;case DIFF_DELETE:b[g]='<del style=\"background:#ffe6e6;\">'+l+\"</del>\";break;case DIFF_EQUAL:b[g]=\"<span>\"+l+\"</span>\"}}return b.join(\"\")};\n" +
"diff_match_patch.prototype.diff_text1=function(a){for(var b=[],c=0;c<a.length;c++)a[c][0]!==DIFF_INSERT&&(b[c]=a[c][1]);return b.join(\"\")};diff_match_patch.prototype.diff_text2=function(a){for(var b=[],c=0;c<a.length;c++)a[c][0]!==DIFF_DELETE&&(b[c]=a[c][1]);return b.join(\"\")};\n" +
"diff_match_patch.prototype.diff_levenshtein=function(a){for(var b=0,c=0,d=0,e=0;e<a.length;e++){var f=a[e][1];switch(a[e][0]){case DIFF_INSERT:c+=f.length;break;case DIFF_DELETE:d+=f.length;break;case DIFF_EQUAL:b+=Math.max(c,d),d=c=0}}return b+=Math.max(c,d)};\n" +
"diff_match_patch.prototype.diff_toDelta=function(a){for(var b=[],c=0;c<a.length;c++)switch(a[c][0]){case DIFF_INSERT:b[c]=\"+\"+encodeURI(a[c][1]);break;case DIFF_DELETE:b[c]=\"-\"+a[c][1].length;break;case DIFF_EQUAL:b[c]=\"=\"+a[c][1].length}return b.join(\"\\t\").replace(/%20/g,\" \")};\n" +
"diff_match_patch.prototype.diff_fromDelta=function(a,b){for(var c=[],d=0,e=0,f=b.split(/\\t/g),g=0;g<f.length;g++){var h=f[g].substring(1);switch(f[g].charAt(0)){case \"+\":try{c[d++]=[DIFF_INSERT,decodeURI(h)]}catch(k){throw Error(\"Illegal escape in diff_fromDelta: \"+h);}break;case \"-\":case \"=\":var l=parseInt(h,10);if(isNaN(l)||0>l)throw Error(\"Invalid number in diff_fromDelta: \"+h);h=a.substring(e,e+=l);\"=\"==f[g].charAt(0)?c[d++]=[DIFF_EQUAL,h]:c[d++]=[DIFF_DELETE,h];break;default:if(f[g])throw Error(\"Invalid diff operation in diff_fromDelta: \"+\n" +
"f[g]);}}if(e!=a.length)throw Error(\"Delta length (\"+e+\") does not equal source text length (\"+a.length+\").\");return c};diff_match_patch.prototype.match_main=function(a,b,c){if(null==a||null==b||null==c)throw Error(\"Null input. (match_main)\");c=Math.max(0,Math.min(c,a.length));return a==b?0:a.length?a.substring(c,c+b.length)==b?c:this.match_bitap_(a,b,c):-1};\n" +
"diff_match_patch.prototype.match_bitap_=function(a,b,c){function d(a,d){var e=a/b.length,g=Math.abs(c-d);return f.Match_Distance?e+g/f.Match_Distance:g?1:e}if(b.length>this.Match_MaxBits)throw Error(\"Pattern too long for this browser.\");var e=this.match_alphabet_(b),f=this,g=this.Match_Threshold,h=a.indexOf(b,c);-1!=h&&(g=Math.min(d(0,h),g),h=a.lastIndexOf(b,c+b.length),-1!=h&&(g=Math.min(d(0,h),g)));var l=1<<b.length-1;h=-1;for(var k,m,p=b.length+a.length,x,w=0;w<b.length;w++){k=0;for(m=p;k<m;)d(w,\n" +
"c+m)<=g?k=m:p=m,m=Math.floor((p-k)/2+k);p=m;k=Math.max(1,c-m+1);var q=Math.min(c+m,a.length)+b.length;m=Array(q+2);for(m[q+1]=(1<<w)-1;q>=k;q--){var t=e[a.charAt(q-1)];m[q]=0===w?(m[q+1]<<1|1)&t:(m[q+1]<<1|1)&t|(x[q+1]|x[q])<<1|1|x[q+1];if(m[q]&l&&(t=d(w,q-1),t<=g))if(g=t,h=q-1,h>c)k=Math.max(1,2*c-h);else break}if(d(w+1,c)>g)break;x=m}return h};\n" +
"diff_match_patch.prototype.match_alphabet_=function(a){for(var b={},c=0;c<a.length;c++)b[a.charAt(c)]=0;for(c=0;c<a.length;c++)b[a.charAt(c)]|=1<<a.length-c-1;return b};\n" +
"diff_match_patch.prototype.patch_addContext_=function(a,b){if(0!=b.length){for(var c=b.substring(a.start2,a.start2+a.length1),d=0;b.indexOf(c)!=b.lastIndexOf(c)&&c.length<this.Match_MaxBits-this.Patch_Margin-this.Patch_Margin;)d+=this.Patch_Margin,c=b.substring(a.start2-d,a.start2+a.length1+d);d+=this.Patch_Margin;(c=b.substring(a.start2-d,a.start2))&&a.diffs.unshift([DIFF_EQUAL,c]);(d=b.substring(a.start2+a.length1,a.start2+a.length1+d))&&a.diffs.push([DIFF_EQUAL,d]);a.start1-=c.length;a.start2-=\n" +
"c.length;a.length1+=c.length+d.length;a.length2+=c.length+d.length}};\n" +
"diff_match_patch.prototype.patch_make=function(a,b,c){if(\"string\"==typeof a&&\"string\"==typeof b&&\"undefined\"==typeof c){var d=a;b=this.diff_main(d,b,!0);2<b.length&&(this.diff_cleanupSemantic(b),this.diff_cleanupEfficiency(b))}else if(a&&\"object\"==typeof a&&\"undefined\"==typeof b&&\"undefined\"==typeof c)b=a,d=this.diff_text1(b);else if(\"string\"==typeof a&&b&&\"object\"==typeof b&&\"undefined\"==typeof c)d=a;else if(\"string\"==typeof a&&\"string\"==typeof b&&c&&\"object\"==typeof c)d=a,b=c;else throw Error(\"Unknown call format to patch_make.\");\n" +
"if(0===b.length)return[];c=[];a=new diff_match_patch.patch_obj;for(var e=0,f=0,g=0,h=d,l=0;l<b.length;l++){var k=b[l][0],m=b[l][1];e||k===DIFF_EQUAL||(a.start1=f,a.start2=g);switch(k){case DIFF_INSERT:a.diffs[e++]=b[l];a.length2+=m.length;d=d.substring(0,g)+m+d.substring(g);break;case DIFF_DELETE:a.length1+=m.length;a.diffs[e++]=b[l];d=d.substring(0,g)+d.substring(g+m.length);break;case DIFF_EQUAL:m.length<=2*this.Patch_Margin&&e&&b.length!=l+1?(a.diffs[e++]=b[l],a.length1+=m.length,a.length2+=m.length):\n" +
"m.length>=2*this.Patch_Margin&&e&&(this.patch_addContext_(a,h),c.push(a),a=new diff_match_patch.patch_obj,e=0,h=d,f=g)}k!==DIFF_INSERT&&(f+=m.length);k!==DIFF_DELETE&&(g+=m.length)}e&&(this.patch_addContext_(a,h),c.push(a));return c};\n" +
"diff_match_patch.prototype.patch_deepCopy=function(a){for(var b=[],c=0;c<a.length;c++){var d=a[c],e=new diff_match_patch.patch_obj;e.diffs=[];for(var f=0;f<d.diffs.length;f++)e.diffs[f]=d.diffs[f].slice();e.start1=d.start1;e.start2=d.start2;e.length1=d.length1;e.length2=d.length2;b[c]=e}return b};\n" +
"diff_match_patch.prototype.patch_apply=function(a,b){if(0==a.length)return[b,[]];a=this.patch_deepCopy(a);var c=this.patch_addPadding(a);b=c+b+c;this.patch_splitMax(a);for(var d=0,e=[],f=0;f<a.length;f++){var g=a[f].start2+d,h=this.diff_text1(a[f].diffs),l=-1;if(h.length>this.Match_MaxBits){var k=this.match_main(b,h.substring(0,this.Match_MaxBits),g);-1!=k&&(l=this.match_main(b,h.substring(h.length-this.Match_MaxBits),g+h.length-this.Match_MaxBits),-1==l||k>=l)&&(k=-1)}else k=this.match_main(b,h,\n" +
"g);if(-1==k)e[f]=!1,d-=a[f].length2-a[f].length1;else if(e[f]=!0,d=k-g,g=-1==l?b.substring(k,k+h.length):b.substring(k,l+this.Match_MaxBits),h==g)b=b.substring(0,k)+this.diff_text2(a[f].diffs)+b.substring(k+h.length);else if(g=this.diff_main(h,g,!1),h.length>this.Match_MaxBits&&this.diff_levenshtein(g)/h.length>this.Patch_DeleteThreshold)e[f]=!1;else{this.diff_cleanupSemanticLossless(g);h=0;var m;for(l=0;l<a[f].diffs.length;l++){var p=a[f].diffs[l];p[0]!==DIFF_EQUAL&&(m=this.diff_xIndex(g,h));p[0]===\n" +
"DIFF_INSERT?b=b.substring(0,k+m)+p[1]+b.substring(k+m):p[0]===DIFF_DELETE&&(b=b.substring(0,k+m)+b.substring(k+this.diff_xIndex(g,h+p[1].length)));p[0]!==DIFF_DELETE&&(h+=p[1].length)}}}b=b.substring(c.length,b.length-c.length);return[b,e]};\n" +
"diff_match_patch.prototype.patch_addPadding=function(a){for(var b=this.Patch_Margin,c=\"\",d=1;d<=b;d++)c+=String.fromCharCode(d);for(d=0;d<a.length;d++)a[d].start1+=b,a[d].start2+=b;d=a[0];var e=d.diffs;if(0==e.length||e[0][0]!=DIFF_EQUAL)e.unshift([DIFF_EQUAL,c]),d.start1-=b,d.start2-=b,d.length1+=b,d.length2+=b;else if(b>e[0][1].length){var f=b-e[0][1].length;e[0][1]=c.substring(e[0][1].length)+e[0][1];d.start1-=f;d.start2-=f;d.length1+=f;d.length2+=f}d=a[a.length-1];e=d.diffs;0==e.length||e[e.length-\n" +
"1][0]!=DIFF_EQUAL?(e.push([DIFF_EQUAL,c]),d.length1+=b,d.length2+=b):b>e[e.length-1][1].length&&(f=b-e[e.length-1][1].length,e[e.length-1][1]+=c.substring(0,f),d.length1+=f,d.length2+=f);return c};\n" +
"diff_match_patch.prototype.patch_splitMax=function(a){for(var b=this.Match_MaxBits,c=0;c<a.length;c++)if(!(a[c].length1<=b)){var d=a[c];a.splice(c--,1);for(var e=d.start1,f=d.start2,g=\"\";0!==d.diffs.length;){var h=new diff_match_patch.patch_obj,l=!0;h.start1=e-g.length;h.start2=f-g.length;\"\"!==g&&(h.length1=h.length2=g.length,h.diffs.push([DIFF_EQUAL,g]));for(;0!==d.diffs.length&&h.length1<b-this.Patch_Margin;){g=d.diffs[0][0];var k=d.diffs[0][1];g===DIFF_INSERT?(h.length2+=k.length,f+=k.length,h.diffs.push(d.diffs.shift()),\n" +
"l=!1):g===DIFF_DELETE&&1==h.diffs.length&&h.diffs[0][0]==DIFF_EQUAL&&k.length>2*b?(h.length1+=k.length,e+=k.length,l=!1,h.diffs.push([g,k]),d.diffs.shift()):(k=k.substring(0,b-h.length1-this.Patch_Margin),h.length1+=k.length,e+=k.length,g===DIFF_EQUAL?(h.length2+=k.length,f+=k.length):l=!1,h.diffs.push([g,k]),k==d.diffs[0][1]?d.diffs.shift():d.diffs[0][1]=d.diffs[0][1].substring(k.length))}g=this.diff_text2(h.diffs);g=g.substring(g.length-this.Patch_Margin);k=this.diff_text1(d.diffs).substring(0,\n" +
"this.Patch_Margin);\"\"!==k&&(h.length1+=k.length,h.length2+=k.length,0!==h.diffs.length&&h.diffs[h.diffs.length-1][0]===DIFF_EQUAL?h.diffs[h.diffs.length-1][1]+=k:h.diffs.push([DIFF_EQUAL,k]));l||a.splice(++c,0,h)}}};diff_match_patch.prototype.patch_toText=function(a){for(var b=[],c=0;c<a.length;c++)b[c]=a[c];return b.join(\"\")};\n" +
"diff_match_patch.prototype.patch_fromText=function(a){var b=[];if(!a)return b;a=a.split(\"\\n\");for(var c=0,d=/^@@ -(\\d+),?(\\d*) \\+(\\d+),?(\\d*) @@$/;c<a.length;){var e=a[c].match(d);if(!e)throw Error(\"Invalid patch string: \"+a[c]);var f=new diff_match_patch.patch_obj;b.push(f);f.start1=parseInt(e[1],10);\"\"===e[2]?(f.start1--,f.length1=1):\"0\"==e[2]?f.length1=0:(f.start1--,f.length1=parseInt(e[2],10));f.start2=parseInt(e[3],10);\"\"===e[4]?(f.start2--,f.length2=1):\"0\"==e[4]?f.length2=0:(f.start2--,f.length2=\n" +
"parseInt(e[4],10));for(c++;c<a.length;){e=a[c].charAt(0);try{var g=decodeURI(a[c].substring(1))}catch(h){throw Error(\"Illegal escape in patch_fromText: \"+g);}if(\"-\"==e)f.diffs.push([DIFF_DELETE,g]);else if(\"+\"==e)f.diffs.push([DIFF_INSERT,g]);else if(\" \"==e)f.diffs.push([DIFF_EQUAL,g]);else if(\"@\"==e)break;else if(\"\"!==e)throw Error('Invalid patch mode \"'+e+'\" in: '+g);c++}}return b};diff_match_patch.patch_obj=function(){this.diffs=[];this.start2=this.start1=null;this.length2=this.length1=0};\n" +
"diff_match_patch.patch_obj.prototype.toString=function(){for(var a=[\"@@ -\"+(0===this.length1?this.start1+\",0\":1==this.length1?this.start1+1:this.start1+1+\",\"+this.length1)+\" +\"+(0===this.length2?this.start2+\",0\":1==this.length2?this.start2+1:this.start2+1+\",\"+this.length2)+\" @@\\n\"],b,c=0;c<this.diffs.length;c++){switch(this.diffs[c][0]){case DIFF_INSERT:b=\"+\";break;case DIFF_DELETE:b=\"-\";break;case DIFF_EQUAL:b=\" \"}a[c+1]=b+encodeURI(this.diffs[c][1])+\"\\n\"}return a.join(\"\").replace(/%20/g,\" \")};\n" +
"this.diff_match_patch=diff_match_patch;this.DIFF_DELETE=DIFF_DELETE;this.DIFF_INSERT=DIFF_INSERT;this.DIFF_EQUAL=DIFF_EQUAL;\n", dirJs+"/diff_match_patch.js");
                
                
                try {
                    File css=new File(dirJs+"/bootstrap.min.css");
                    File js=new File(dirJs+"/bootstrap.min.js");
                    File jq=new File(dirJs+"/jquery.min.js"); 
                    File po=new File(dirJs+"/popper.min.js"); 
                    File cu=new File(dirJs+"/cuerpo.js"); 
                    File na=new File(dirJs+"/nav.js"); 
                    File ht=new File(dir+"/index.html"); 
                    streamToFile(this.getClass().getResourceAsStream("bootstrap.min.txt"),css.getPath());
                    streamToFile(this.getClass().getResourceAsStream("bootstrap.min.js"),js.getPath());
                    streamToFile(this.getClass().getResourceAsStream("jquery.min.js"),jq.getPath()); 
                    streamToFile(this.getClass().getResourceAsStream("popper.min.js"),po.getPath()); 
                    streamToFile(this.getClass().getResourceAsStream("cuerpo.js"),cu.getPath()); 
                    streamToFile(this.getClass().getResourceAsStream("nav.js"),na.getPath()); 
                    streamToFile(this.getClass().getResourceAsStream("index.html"),ht.getPath()); 

                } catch (IOException ex) {
                    Logger.getLogger(VentanaPrincipal.class.getName()).log(Level.SEVERE, null, ex);
                } 
                String entrada = dirJs+"/entrada.js";
                generarDiffHtml(fileIn.getAbsolutePath(),fileOut.getAbsolutePath(),entrada);
                seletorArchivosIn.updateUI();
                File htmlFile = new File(dir+"/index.html");
                try {
                    Desktop.getDesktop().browse(htmlFile.toURI());
                } catch (IOException ex) {
                    Logger.getLogger(VentanaPrincipal.class.getName()).log(Level.SEVERE, null, ex);
                } 
            }

        }
    }//GEN-LAST:event_botonDiferenciaActionPerformed

    private void botonCorregirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonCorregirActionPerformed
        // TODO add your handling code here:
        File fileIn = this.seletorArchivosIn.getSelectedFile();
        if (fileIn != null){
            JPanel panel = new JPanel();
            JFileChooser fc = new JFileChooser();
            panel.add(fc);
            fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
            fc.setApproveButtonText("Corregir");
            panel.setVisible(true);
            int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
            if (seleccion == JFileChooser.APPROVE_OPTION){
                File fileOut = fc.getSelectedFile();
                if (fileOut != null){ 
                    String resultado=ejecutar(ejecutablePath+" -op=r -in=\""+fileIn+"\" -out=\""+convertirHam(fileOut.getAbsolutePath()))+"\"";
                    if (resultado.length() >1){
                        JOptionPane.showMessageDialog(new JFrame(),resultado,"Corregir error", JOptionPane.ERROR_MESSAGE);
                    }
                    seletorArchivosIn.updateUI();
                }
            }

        }
    }//GEN-LAST:event_botonCorregirActionPerformed

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
                    String resultado=ejecutar(ejecutablePath+" -op=i -in=\""+fileIn+"\" -out=\""+convertirHam(fileOut.getAbsolutePath())+"\"");
                    if (resultado.length() >1){
                        JOptionPane.showMessageDialog(new JFrame(),resultado,"Introducir error", JOptionPane.ERROR_MESSAGE);
                    }
                    seletorArchivosIn.updateUI();
                }
            }

        }
    }//GEN-LAST:event_botonDañarActionPerformed

    private void botonComprobarActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonComprobarActionPerformed
        File fileIn = this.seletorArchivosIn.getSelectedFile();
        if (fileIn != null){ 
            String resultado=ejecutar(ejecutablePath+" -op=e -in=\""+fileIn+"\"");
            String[] arr=resultado.split("\n");
            if (arr.length > 1){ 
                String mensajeErrores = getStringErrores(arr);
                JOptionPane.showMessageDialog(new JFrame(), "Contiene error: "+ mensajeErrores, "Comprobar error", JOptionPane.ERROR_MESSAGE);
            }else{
                JOptionPane.showMessageDialog(new JFrame(),"No contiene error "+fileIn.getName() , "Comprobar error", JOptionPane.INFORMATION_MESSAGE);
            }
        }
    }//GEN-LAST:event_botonComprobarActionPerformed

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
                String resultado=ejecutar(ejecutablePath+" -op=dp -in=\""+fileIn+"\" -out=\""+fileOut+"\"");
                if (resultado.length() >1){
                    JOptionPane.showMessageDialog(new JFrame(),"<html><b>Archivo entrada: </b>"+fileIn+"<br><b>Archivo salida: </b>"+fileOut+"<br><b>Tardo: </b>"+resultado,"Desproteger Archivo", JOptionPane.INFORMATION_MESSAGE);
                }
                seletorArchivosIn.updateUI();
            }
        }
    }//GEN-LAST:event_botonDesprotegerActionPerformed

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
                String resultado=ejecutar(ejecutablePath+" -op=p -in=\""+fileIn+"\" -out=\""+convertirHam(fileOut.getAbsolutePath())+"\" -cod="+codificacion);
                if (resultado.length() >1){
                    JOptionPane.showMessageDialog(new JFrame(),"<html><b>Archivo entrada: </b>"+fileIn+"<br><b>Archivo salida: </b>"+fileOut+"<br><b>Tardo: </b>"+resultado ,"Proteger Archivo", JOptionPane.INFORMATION_MESSAGE);
                }
                seletorArchivosIn.updateUI();
            }
        }
    }//GEN-LAST:event_botonProtegerActionPerformed

    private void botonDescomprimirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonDescomprimirActionPerformed
       File fileIn = this.seletorArchivosIn.getSelectedFile();

        String pathOut = "";
        JPanel panel = new JPanel();
        JFileChooser fc = new JFileChooser();
        panel.add(fc);
        fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
        fc.setApproveButtonText("Descomprimir");
        panel.setVisible(true);
        int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
        if (seleccion == JFileChooser.APPROVE_OPTION){
            File fileOut = fc.getSelectedFile();
            if (fileOut != null){
                String resultado=ejecutar(ejecutablePath+" -op=d -in=\""+fileIn+"\" -out=\""+fileOut+"\""); 
                JOptionPane.showMessageDialog(new JFrame(),resultado , "Descomprimir", JOptionPane.INFORMATION_MESSAGE);
                seletorArchivosIn.updateUI();
            }
        }
    }//GEN-LAST:event_botonDescomprimirActionPerformed

    private void botonComprimirActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_botonComprimirActionPerformed
         File fileIn = this.seletorArchivosIn.getSelectedFile();

        String pathOut = "";
        JPanel panel = new JPanel();
        JFileChooser fc = new JFileChooser();
        panel.add(fc);
        fc.setFileSelectionMode(JFileChooser.FILES_ONLY);
        fc.setApproveButtonText("Comprimir");
        panel.setVisible(true);
        int seleccion =fc.showSaveDialog(jPanel1);// sacar jPanell si no anda
        if (seleccion == JFileChooser.APPROVE_OPTION){
            File fileOut = fc.getSelectedFile();
            if (fileOut != null){ 
                String resultado=ejecutar(ejecutablePath+" -op=c -in=\""+fileIn+"\" -out=\""+fileOut+"\""); 
                JOptionPane.showMessageDialog(new JFrame(),resultado , "Comprimir", JOptionPane.INFORMATION_MESSAGE);
                seletorArchivosIn.updateUI();
            }
        }
    }//GEN-LAST:event_botonComprimirActionPerformed

    private void stringToFile(String entrada, String fileOut){
         try { 
            FileWriter fileWriter = new FileWriter(fileOut);
            BufferedWriter bufferedWriter =
                new BufferedWriter(fileWriter); 
             String[] split = entrada.split("\n"); 
             for (int i=0;i< split.length;i++){
                bufferedWriter.write(split[i]+"\n");   
             }
            bufferedWriter.close();
        }
        catch(IOException ex) {
            System.out.println(
                "Error writing to file '"
                + fileOut + "'"); 
        }
    }
    
    private String fileToString(String fileName) { 
     String line = null; 
     StringBuilder res=new StringBuilder(""); 
        try { 
            FileReader fileReader = 
                new FileReader(fileName);
 
            BufferedReader bufferedReader = 
                new BufferedReader(fileReader);

            while((line = bufferedReader.readLine()) != null) {
                res.append(line+"\n");
            }   
 
            bufferedReader.close();         
        }
        catch(FileNotFoundException ex) {
            System.out.println(
                "Unable to open file '" + 
                fileName + "'");                
        }
        catch(IOException ex) {
            System.out.println(
                "Error reading file '" 
                + fileName + "'");          
        } 
        return res.toString();

}
    private String generarHtml(String fileIn, String fileOut){
        String text1= fileToString(fileIn);
        text1=text1.replace("`", "'");
        String text2= fileToString(fileOut);
        text2=text2.replace("`", "'");  
        return 
"            var dmp = new diff_match_patch();" + 
"            var diff = dmp.diff_main( `"+text1+"`,  `"+text2+"`);\n";
    }
    private void generarDiffHtml( String fileIn, String fileOut, String html){
       stringToFile(generarHtml(fileIn,fileOut),html);
    }
    private void CopyFile(String fileIn, String fileOut) throws IOException {
         InputStream is = null;
        OutputStream os = null;
        try {
             Logger.getLogger(VentanaPrincipal.class.getName()).log(Level.INFO, fileIn+fileOut); 
            is = new FileInputStream(fileIn);
            os = new FileOutputStream(fileOut);
            byte[] buffer = new byte[1024];
            int length;
            while ((length = is.read(buffer)) > 0) {
                os.write(buffer, 0, length);
            }
        } finally {
            is.close();
            os.close();
        }

    
    }    
    /**
     * @param args the command line arguments
     */
    public static void main(String args[]) {


        /* Create and display the form */
        java.awt.EventQueue.invokeLater(new Runnable() {
            public void run() {
                new VentanaPrincipal().setVisible(true);
            }
        });
        
    }

    // Variables declaration - do not modify//GEN-BEGIN:variables
    private javax.swing.JMenuBar BarraMenu;
    private javax.swing.JFrame Otro;
    private javax.swing.JMenuItem botonAcerca;
    private javax.swing.JButton botonComprimir;
    private javax.swing.JButton botonComprobar;
    private javax.swing.JButton botonCorregir;
    private javax.swing.JButton botonDañar;
    private javax.swing.JButton botonDescomprimir;
    private javax.swing.JButton botonDesproteger;
    private javax.swing.JButton botonDiferencia;
    private javax.swing.JMenuItem botonEditor;
    private javax.swing.JMenu botonPreferencias;
    private javax.swing.JButton botonProteger;
    private javax.swing.JMenuItem botonSalir;
    private javax.swing.Box.Filler filler1;
    private javax.swing.Box.Filler filler2;
    private javax.swing.JPanel jPanel1;
    private javax.swing.JPanel jPanel2;
    private javax.swing.JPopupMenu.Separator jSeparator1;
    private javax.swing.JPopupMenu.Separator jSeparator2;
    private javax.swing.JMenu menuCodificacion;
    private javax.swing.JFileChooser seletorArchivosIn;
    // End of variables declaration//GEN-END:variables

  
}
