import React from 'react';
import { IconButton } from '../../atoms/IconButton/IconButton';
import type { IconButtonProps } from '../../atoms/IconButton/IconButton';
import styles from './SideMenuButton.module.css';

interface SideMenuButtonProps extends IconButtonProps {
  animationDuration?: number;
}

export const SideMenuButton = ({
  icon,
  text,
  backgroundColor = '#ffffff',
  hoverColor = '#f5f5f5',
  animationDuration = 300,
  onClick,
}: SideMenuButtonProps) => {
  return (
    <div
      className={styles.container}
      style={
        {
          '--animation-duration': `${animationDuration}ms`,
        } as React.CSSProperties
      }
    >
      <IconButton
        icon={icon}
        text={text}
        backgroundColor={backgroundColor}
        hoverColor={hoverColor}
        onClick={onClick}
      />
    </div>
  );
};
