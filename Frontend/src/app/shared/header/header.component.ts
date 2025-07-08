import { Component } from '@angular/core';
import { MenuModule } from 'primeng/menu';
import { AuthComponent } from '../../core/auth/auth.component';

@Component({
  selector: 'app-header',
  templateUrl: 'header.component.html',
  styleUrl: 'header.component.css',
  imports: [MenuModule, AuthComponent],
})
export class HeaderComponent {
  appName = 'traveloka';

  productLinks = [
    { name: 'Khách sạn', url: '#' },
    { name: 'Vé máy bay', url: '#' },
    { name: 'Vé xe khách', url: '#' },
    { name: 'Đưa đón sân bay', url: '#' },
    { name: 'Cho thuê xe', url: '#' },
    { name: 'Hoạt động & Vui chơi', url: '#' },
  ];

  countryLanguageCurrency = [
    {
      flag: 'fi fi-id',
      country: 'Indonesia',
      language: 'Bahasa Indonesia',
      currency: { currency: 'IDR', name: 'Rupiah' },
    },
    {
      flag: 'fi fi-id',
      country: 'Indonesia',
      language: 'English',
      currency: { currency: 'IDR', name: 'Rupiah' },
    },
    {
      flag: 'fi fi-un',
      country: 'Global',
      language: 'English',
      currency: [
        { currency: 'EUR', name: 'Euro' },
        { currency: 'USD', name: 'US Dollar' },
        { currency: 'CNY', name: 'Renminbi (Chinese) Yuan' },
        { currency: 'CAD', name: 'Canadian Dollar' },
        { currency: 'HKD', name: 'Hong Kong Dollar' },
        { currency: 'GBP', name: 'British Pound Sterling' },
        { currency: 'CHF', name: 'Swiss Franc' },
        { currency: 'NZD', name: 'New Zealand Dollar' },
        { currency: 'TWD', name: 'New Taiwan Dollar' },
        { currency: 'BRL', name: 'Brazilian Real' },
        { currency: 'SAR', name: 'Saudi Arabian Riyal' },
        { currency: 'INR', name: 'Indian Rupee' },
        { currency: 'QAR', name: 'Qatari Riyal' },
        { currency: 'AED', name: 'United Arab Emirates Dirham' },
      ],
    },
    {
      class: 'fi fi-my',
      country: 'Malaysia',
      language: 'Bahasa Malaysia',
      currency: { currency: 'MYR', name: 'Malaysian Ringgit' },
    },
    {
      class: 'fi fi-my',
      country: 'Malaysia',
      language: 'English',
      currency: { currency: 'MYR', name: 'Malaysian Ringgit' },
    },
    {
      flag: 'fi fi-ph',
      country: 'Philippines',
      language: 'English',
      currency: { currency: 'PHP', name: 'Philippine Peso' },
    },
    {
      flag: 'fi fi-sg',
      country: 'Singapore',
      language: 'English',
      currency: { currency: 'SGD', name: 'Singapore Dollar' },
    },
    {
      flag: 'fi fi-th',
      country: 'ไทย',
      language: 'ภาษาไทย',
      currency: { currency: 'THB', name: 'Thai Baht' },
    },
    {
      flag: 'fi fi-th',
      country: 'Thailand',
      language: 'English',
      currency: { currency: 'THB', name: 'Thai Baht' },
    },
    {
      flag: 'fi fi-vn',
      country: 'Việt Nam',
      language: 'Tiếng Việt',
      currency: { currency: 'VND', name: 'Vietnamese Dong' },
    },
    {
      flag: 'fi fi-vn',
      country: 'Vietnam',
      language: 'English',
      currency: { currency: 'VND', name: 'Vietnamese Dong' },
    },
    {
      flag: 'fi fi-au',
      country: 'Australia',
      language: 'English',
      currency: { currency: 'AUD', name: 'Australian Dollar' },
    },
    {
      flag: 'fi fi-kr',
      country: '대한민국',
      language: '한국어',
      currency: { currency: 'KRW', name: 'South Korean Won' },
    },
    {
      flag: 'fi fi-kr',
      country: 'Korea',
      language: 'English',
      currency: { currency: 'KRW', name: 'South Korean Won' },
    },
    {
      flag: 'fi fi-jp',
      country: '日本',
      language: '日本語',
      currency: { currency: 'JPY', name: 'Japanese Yen' },
    },
    {
      flag: 'fi fi-jp',
      country: 'Japan',
      language: 'English',
      currency: { currency: 'JPY', name: 'Japanese Yen' },
    },
  ];

  flag = 'fi fi-vn';
}
