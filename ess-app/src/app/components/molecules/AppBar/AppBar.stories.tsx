import type { Meta, StoryObj } from '@storybook/react';
import { AppBar } from './AppBar';
import { Icon } from '../../atoms/Icon/Icon';

const meta: Meta<typeof AppBar> = {
  title: 'molecules/AppBar',
  component: AppBar,
};

export default meta;
type Story = StoryObj<typeof AppBar>;

export const Default: Story = {
  args: {
    icon: <Icon name='menu' />,
    text: 'アプリケーション名',
  },
};

export const CustomColor: Story = {
  args: {
    icon: <Icon name='menu' />,
    text: 'カスタムカラー',
    color: '#ffffff',
    backgroundColor: '#3f51b5',
  },
};

export const WithNavigation: Story = {
  args: {
    icon: <Icon name='home' />,
    text: 'ホームに戻る',
    iconHref: '/',
  },
};
