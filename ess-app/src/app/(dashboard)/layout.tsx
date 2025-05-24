import { ReactNode } from 'react';
import { AppBar } from '@/app/components/molecules/AppBar/AppBar';
import { SideMenuButton } from '@/app/components/molecules/SideMenuButton/SideMenuButton';
import styles from './layout.module.css';

export default function DashboardLayout({ children }: { children: ReactNode }) {
  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <AppBar
          icon='menu'
          text='後で名前をつける'
          backgroundColor='#C8E4F4'
          color='white'
        />
      </div>
      <div className={styles.sidebar}>
        <SideMenuButton icon='home' text='Home' />
        <SideMenuButton icon='home' text='Home' />
        <SideMenuButton icon='home' text='Home' />
        {/* 他のメニューボタンを追加可能 */}
      </div>
      <main className={styles.main}>{children}</main>
    </div>
  );
}
