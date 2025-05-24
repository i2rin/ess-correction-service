import React from 'react';
import styles from './IconButton.module.css';

export interface IconButtonProps {
  icon: React.ReactNode;
  text: string;
  backgroundColor?: string;
  hoverColor?: string;
  onClick?: () => void;
}

export const IconButton = ({
  icon,
  text,
  backgroundColor = '#ffffff',
  hoverColor = '#f5f5f5',
  onClick,
}: IconButtonProps) => {
  return (
    <button
      className={styles.button}
      style={
        {
          '--bg-color': backgroundColor,
          '--hover-color': hoverColor,
        } as React.CSSProperties
      }
      onClick={onClick}
    >
      <span className={styles.icon}>{icon}</span>
      <span className={styles.text}>{text}</span>
    </button>
  );
};
