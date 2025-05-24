import React from 'react';
import Link from 'next/link';
import styles from './AppBar.module.css';

interface AppBarProps {
  icon: React.ReactNode;
  text: string;
  color?: string;
  backgroundColor?: string;
  iconHref?: string;
  onIconClick?: () => void;
}

export const AppBar = ({
  icon,
  text,
  color = '#000000',
  backgroundColor = '#ffffff',
  iconHref,
  onIconClick,
}: AppBarProps) => {
  return (
    <div className={styles.appBar} style={{ backgroundColor }}>
      <div className={styles.content}>
        {iconHref ? (
          <Link
            href={iconHref}
            className={styles.icon}
            style={{ color }}
            onClick={onIconClick}
          >
            {icon}
          </Link>
        ) : (
          <div className={styles.icon} style={{ color }} onClick={onIconClick}>
            {icon}
          </div>
        )}
        <span className={styles.text} style={{ color }}>
          {text}
        </span>
      </div>
    </div>
  );
};
