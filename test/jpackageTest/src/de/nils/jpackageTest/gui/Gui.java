package de.nils.jpackageTest.gui;

import java.awt.BorderLayout;

import javax.swing.BorderFactory;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.SwingConstants;

public class Gui
{
	private int count = 0;
	
	public Gui() 
	{
		JFrame frame = new JFrame("jpackageTest");
		frame.setLayout(new BorderLayout());
		
		JLabel label = new JLabel(count + "x Clicked", SwingConstants.CENTER);
		label.setBorder(BorderFactory.createEmptyBorder(10, 10, 10, 10));
		
		JButton button = new JButton("Click me");
		button.addActionListener(l ->
		{
			count++;
			label.setText(count + "x Clicked");
		});
		
		frame.add(button, BorderLayout.CENTER);
		frame.add(label, BorderLayout.SOUTH);
		
		frame.setSize(300, 300);
		frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		frame.setLocationRelativeTo(null);
		frame.setVisible(true);
	}
}
