# Uy vazifasi: Go RESTful API-da Foydalanuvchilarni Qo'shish Va Dinamik CORSni Cheklash

## Maqsad
Bu vazifada, foydalanuvchilar uchun autentifikatsiya qilish funksiyasini `RESTful API`'ga qo'shib, `CORS` operatsiyalarini faqat ma'lum foydalanuvchilarga ruxsat berishingiz kerak bo'ladi. Shuningdek, `Redis` orqali saqlanadigan ishonchli `origin`lar va foydalanuvchilar ma'lumotlarini boshqarasiz.

## Talablar
1. **RESTful API bilan foydalanuvchilarni boshqarish:**:
    - Yangi foydalanuvchi ro'yxatdan o'tkazish: `POST /register`
    - Foydalanuvchi login qilish: `POST /login`
    - Autentifikatsiya qiluvchi tokenni yaratish: `JWT` orqali foydalanuvchilar uchun tokenni yarating.

2. **Foydalanuvchilarni CORS operatsiyalariga cheklash:**:
    - Faqatgina autentifikatsiyalangan foydalanuvchilar ishonchli `origin`lar qo'shish va o'chirish imkoniyatiga ega bo'lishlari kerak.
    - Autentifikatsiyalangan foydalanuvchilar token orqali har bir so'rov davomida tekshirilishi kerak.
    
3. **Dinamik CORSni Tekshirish:**:
    - Middleware yordamida foydalanuvchi autentifikatsiyasini qo'shing va `Redis`da saqlangan foydalanuvchilar asosida `CORS`ni boshqaring.    

4. **CRUD Operatsiyalari:**:
    - Ishonchli `origin` larni boshqarish va faqat autentifikatsiyalangan foydalanuvchilar uchun ularni olish, qo'shish va o'chirish imkoniyatini yarating.   
 















