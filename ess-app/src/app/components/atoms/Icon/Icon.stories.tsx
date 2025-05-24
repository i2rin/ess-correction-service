import type { Meta, StoryObj } from '@storybook/react';
import { Icon } from './Icon';

const meta: Meta<typeof Icon> = {
  title: 'atoms/Icon',
  component: Icon,
};

export default meta;
type Story = StoryObj<typeof Icon>;

export const Default: Story = {
  args: {
    name: 'menu',
  },
};

export const WithSize: Story = {
  args: {
    name: 'home',
    size: 32,
  },
};

export const WithColor: Story = {
  args: {
    name: 'home',
    color: '#3f51b5',
  },
};
