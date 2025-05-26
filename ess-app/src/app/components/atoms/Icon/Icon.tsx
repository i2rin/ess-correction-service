import React from 'react';
import styles from './Icon.module.css';

interface IconProps {
  name: string;
  size?: number;
  color?: string;
}

export const Icon = ({
  name,
  size = 24,
  color = 'currentColor',
}: IconProps) => {
  return (
    <span
      className={styles.icon}
      style={{
        fontSize: `${size}px`,
        color: color,
      }}
    >
      {name}
    </span>
  );
};
