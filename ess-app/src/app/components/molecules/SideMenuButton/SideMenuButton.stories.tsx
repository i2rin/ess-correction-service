import type { Meta, StoryObj } from '@storybook/react';
import { SideMenuButton } from './SideMenuButton';
import { FiHome, FiUser, FiSettings } from 'react-icons/fi';

const meta: Meta<typeof SideMenuButton> = {
  title: 'Molecules/SideMenuButton',
  component: SideMenuButton,
  tags: ['autodocs'],
  argTypes: {
    backgroundColor: { control: 'color' },
    hoverColor: { control: 'color' },
    animationDuration: { control: 'number' },
  },
};

export default meta;
type Story = StoryObj<typeof SideMenuButton>;

export const Default: Story = {
  args: {
    icon: <FiHome size={24} />,
    text: 'ホーム',
    backgroundColor: '#ffffff',
    hoverColor: '#f5f5f5',
    animationDuration: 300,
  },
};

export const WithDifferentIcons: Story = {
  render: () => (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
      <SideMenuButton
        icon={<FiHome size={24} />}
        text='ホーム'
        backgroundColor='#ffffff'
        hoverColor='#f5f5f5'
      />
      <SideMenuButton
        icon={<FiUser size={24} />}
        text='プロフィール'
        backgroundColor='#ffffff'
        hoverColor='#f5f5f5'
      />
      <SideMenuButton
        icon={<FiSettings size={24} />}
        text='設定'
        backgroundColor='#ffffff'
        hoverColor='#f5f5f5'
      />
    </div>
  ),
};
