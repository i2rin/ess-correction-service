import type { Meta, StoryObj } from '@storybook/react';
import { IconButton } from './IconButton';
import { FiHome } from 'react-icons/fi';

const meta: Meta<typeof IconButton> = {
  title: 'Atoms/IconButton',
  component: IconButton,
  tags: ['autodocs'],
  argTypes: {
    backgroundColor: { control: 'color' },
    hoverColor: { control: 'color' },
  },
};

export default meta;
type Story = StoryObj<typeof IconButton>;

export const Default: Story = {
  args: {
    icon: <FiHome size={24} />,
    text: 'ホーム',
    backgroundColor: '#ffffff',
    hoverColor: '#f5f5f5',
  },
};

export const Dark: Story = {
  args: {
    icon: <FiHome size={24} />,
    text: 'ダークモード',
    backgroundColor: '#333333',
    hoverColor: '#555555',
  },
};
