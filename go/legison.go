package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(ID_ENFORCER_REGX)
	var str = `
QUỐC HỘI

 ____________

Số: 55/2005/QH11
	

CỘNG HOÀ XÃ HỘI CHỦ NGHĨA VIỆT NAM

Độc lập - Tự do - Hạnh phúc

____________________

Hà Nội , Ngày 29 tháng 11 năm 2005
 

 

LUẬT

PHÒNG, CHỐNG THAM NHŨNG

 

Căn cứ vào Hiến pháp nước Cộng hoà xã hội chủ nghĩa Việt Nam năm 1992 đã được sửa đổi, bổ sung theo Nghị quyết số 51/2001/QH10 ngày 25 tháng 12 năm 2001 của Quốc hội khoá X, kỳ họp thứ 10;

Luật này quy định về phòng, chống tham nhũng.

Chương I



NHỮNG QUY ĐỊNH CHUNG

Điều 1. Phạm vi điều chỉnh

1. Luật này quy định về phòng ngừa, phát hiện, xử lý người có hành vi tham nhũng và trách nhiệm của cơ quan, tổ chức, đơn vị, cá nhân trong phòng, chống tham nhũng.

2. Tham nhũng là hành vi của người có chức vụ, quyền hạn đã lợi dụng chức vụ, quyền hạn đó vì vụ lợi.

3. Người có chức vụ, quyền hạn bao gồm:

a) Cán bộ, công chức, viên chức;

b) Sĩ quan, quân nhân chuyên nghiệp, công nhân quốc phòng trong cơ quan, đơn vị thuộc Quân đội nhân dân; sĩ quan, hạ sĩ quan nghiệp vụ, sĩ quan, hạ sĩ quan chuyên môn - kỹ thuật trong cơ quan, đơn vị thuộc Công an nhân dân;

c) Cán bộ lãnh đạo, quản lý trong doanh nghiệp của Nhà nước; cán bộ lãnh đạo, quản lý là người đại diện phần vốn góp của Nhà nước tại doanh nghiệp;

d) Người được giao thực hiện nhiệm vụ, công vụ có quyền hạn trong khi thực hiện nhiệm vụ, công vụ đó.

Điều 2. Giải thích từ ngữ

Trong Luật này, các từ ngữ dưới đây được hiểu như sau:

1. Tài sản tham nhũng là tài sản có được từ hành vi tham nhũng, tài sản có nguồn gốc từ hành vi tham nhũng.

2. Công khai là việc cơ quan, tổ chức, đơn vị công bố, cung cấp thông tin chính thức về văn bản, hoạt động hoặc về nội dung nhất định.

3. Minh bạch tài sản, thu nhập là việc kê khai tài sản, thu nhập của người có nghĩa vụ kê khai và khi cần thiết được xác minh, kết luận.

4. Nhũng nhiễu là hành vi cửa quyền, hách dịch, gây khó khăn, phiền hà khi thực hiện nhiệm vụ, công vụ.

5. Vụ lợi là lợi ích vật chất, tinh thần mà người có chức vụ, quyền hạn đạt được hoặc có thể đạt được thông qua hành vi tham nhũng.

6. Cơ quan, tổ chức, đơn vị bao gồm cơ quan nhà nước, tổ chức chính trị, tổ chức chính trị - xã hội, đơn vị vũ trang nhân dân, đơn vị sự nghiệp, doanh nghiệp của Nhà nước và cơ quan, tổ chức, đơn vị khác có sử dụng ngân sách, tài sản của Nhà nước.

Điều 3. Các hành vi tham nhũng

1. Tham ô tài sản.

2. Nhận hối lộ.

3. Lạm dụng chức vụ, quyền hạn chiếm đoạt tài sản.

4. Lợi dụng chức vụ, quyền hạn trong khi thi hành nhiệm vụ, công vụ vì vụ lợi.

5. Lạm quyền trong khi thi hành nhiệm vụ, công vụ vì vụ lợi.

6. Lợi dụng chức vụ, quyền hạn gây ảnh hưởng với người khác để trục lợi.

7. Giả mạo trong công tác vì vụ lợi.

8. Đưa hối lộ, môi giới hối lộ được thực hiện bởi người có chức vụ, quyền hạn để giải quyết công việc của cơ quan, tổ chức, đơn vị hoặc địa phương vì vụ lợi.

9. Lợi dụng chức vụ, quyền hạn sử dụng trái phép tài sản của Nhà nước vì vụ lợi.

10. Nhũng nhiễu vì vụ lợi.

11. Không thực hiện nhiệm vụ, công vụ vì vụ lợi.

12. Lợi dụng chức vụ, quyền hạn để bao che cho người có hành vi vi phạm pháp luật vì vụ lợi; cản trở, can thiệp trái pháp luật vào việc kiểm tra, thanh tra, kiểm toán, điều tra, truy tố, xét xử, thi hành án vì vụ lợi.

Điều 4. Nguyên tắc xử lý tham nhũng

1. Mọi hành vi tham nhũng đều phải được phát hiện, ngăn chặn và xử lý kịp thời, nghiêm minh.

2. Người có hành vi tham nhũng ở bất kỳ cương vị, chức vụ nào phải bị xử lý theo quy định của pháp luật.

3. Tài sản tham nhũng phải được thu hồi, tịch thu; người có hành vi tham nhũng gây thiệt hại thì phải bồi thường, bồi hoàn theo quy định của pháp luật.

4. Người có hành vi tham nhũng đã chủ động khai báo trước khi bị phát hiện, tích cực hạn chế thiệt hại do hành vi trái pháp luật của mình gây ra, tự giác nộp lại tài sản tham nhũng thì có thể được xem xét giảm nhẹ hình thức kỷ luật, giảm nhẹ hình phạt hoặc miễn truy cứu trách nhiệm hình sự theo quy định của pháp luật.

5. Việc xử lý tham nhũng phải được thực hiện công khai theo quy định của pháp luật.

6. Người có hành vi tham nhũng đã nghỉ hưu, thôi việc, chuyển công tác vẫn phải bị xử lý về hành vi tham nhũng do mình đã thực hiện.

Điều 5. Trách nhiệm của cơ quan, tổ chức, đơn vị và người có chức vụ, quyền hạn

1. Cơ quan, tổ chức, đơn vị trong phạm vi nhiệm vụ, quyền hạn của mình có trách nhiệm sau đây:

a) Tổ chức thực hiện văn bản quy phạm pháp luật về phòng, chống tham nhũng;

b) Tiếp nhận, xử lý kịp thời báo cáo, tố giác, tố cáo và thông tin khác về hành vi tham nhũng;

c) Bảo vệ quyền và lợi ích hợp pháp của người phát hiện, báo cáo, tố giác, tố cáo hành vi tham nhũng;

d) Chủ động phòng ngừa, phát hiện hành vi tham nhũng; kịp thời cung cấp thông tin, tài liệu và thực hiện yêu cầu của cơ quan, tổ chức, cá nhân có thẩm quyền trong quá trình phát hiện, xử lý người có hành vi tham nhũng.

2. Người đứng đầu cơ quan, tổ chức, đơn vị trong phạm vi nhiệm vụ, quyền hạn của mình có trách nhiệm sau đây:

a) Chỉ đạo việc thực hiện các quy định tại khoản 1 Điều này;

b) Gương mẫu, liêm khiết; định kỳ kiểm điểm việc thực hiện chức trách, nhiệm vụ và trách nhiệm của mình trong việc phòng ngừa, phát hiện hành vi tham nhũng, xử lý người có hành vi tham nhũng;

c) Chịu trách nhiệm khi để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách.

3. Người có chức vụ, quyền hạn có trách nhiệm sau đây:

a) Thực hiện nhiệm vụ, công vụ đúng quy định của pháp luật;

b) Gương mẫu, liêm khiết; chấp hành nghiêm chỉnh quy định của pháp luật về phòng, chống tham nhũng, quy tắc ứng xử, quy tắc đạo đức nghề nghiệp;

c) Kê khai tài sản theo quy định của Luật này và chịu trách nhiệm về tính chính xác, trung thực của việc kê khai đó.

Điều 6. Quyền và nghĩa vụ của công dân trong phòng, chống tham nhũng

Công dân có quyền phát hiện, tố cáo hành vi tham nhũng; có nghĩa vụ hợp tác, giúp đỡ cơ quan, tổ chức, cá nhân có thẩm quyền trong việc phát hiện, xử lý người có hành vi tham nhũng.

Điều 7. Trách nhiệm phối hợp của cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án và của cơ quan, tổ chức, đơn vị hữu quan

Cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án trong phạm vi nhiệm vụ, quyền hạn của mình có trách nhiệm phối hợp với nhau và phối hợp với cơ quan, tổ chức, đơn vị hữu quan trong việc phát hiện hành vi tham nhũng, xử lý người có hành vi tham nhũng và phải chịu trách nhiệm trước pháp luật về kết luận, quyết định của mình trong quá trình thanh tra, kiểm toán, điều tra, truy tố, xét xử vụ việc tham nhũng.

Cơ quan, tổ chức, đơn vị hữu quan có trách nhiệm tạo điều kiện, cộng tác với cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án trong việc phát hiện, xử lý người có hành vi tham nhũng.

Điều 8. Trách nhiệm của Mặt trận Tổ quốc Việt Nam và các tổ chức thành viên

Mặt trận Tổ quốc Việt Nam và các tổ chức thành viên có trách nhiệm động viên nhân dân tham gia tích cực vào việc phòng, chống tham nhũng; phát hiện, kiến nghị cơ quan, tổ chức, cá nhân có thẩm quyền xử lý người có hành vi tham nhũng; giám sát việc thực hiện pháp luật về phòng, chống tham nhũng.

Điều 9. Trách nhiệm của cơ quan báo chí

Cơ quan báo chí có trách nhiệm tham gia vào việc phòng, chống tham nhũng; hợp tác với cơ quan, tổ chức, cá nhân có thẩm quyền trong phòng, chống tham nhũng; khi đưa tin phải bảo đảm chính xác, trung thực, khách quan và phải chịu trách nhiệm về nội dung của thông tin đã đưa.

Điều 10. Các hành vi bị nghiêm cấm

1. Các hành vi quy định tại Điều 3 của Luật này.

2. Đe doạ, trả thù, trù dập người phát hiện, báo cáo, tố giác, tố cáo, cung cấp thông tin về hành vi tham nhũng.

3. Lợi dụng việc tố cáo tham nhũng để vu cáo, vu khống cơ quan, tổ chức, đơn vị, cá nhân khác.

Chương II

PHÒNG NGỪA THAM NHŨNG

Mục 1

CÔNG KHAI, MINH BẠCH TRONG HOẠT ĐỘNG

CỦA CƠ QUAN, TỔ CHỨC, ĐƠN VỊ

Điều 11. Nguyên tắc và nội dung công khai, minh bạch trong hoạt động của cơ quan, tổ chức, đơn vị

1. Chính sách, pháp luật và việc tổ chức thực hiện chính sách, pháp luật phải được công khai, minh bạch, bảo đảm công bằng, dân chủ.

2. Cơ quan, tổ chức, đơn vị phải công khai hoạt động của mình, trừ nội dung thuộc bí mật nhà nước và những nội dung khác theo quy định của Chính phủ.

Điều 12. Hình thức công khai

1. Hình thức công khai bao gồm:

a) Công bố tại cuộc họp của cơ quan, tổ chức, đơn vị;

b) Niêm yết tại trụ sở làm việc của cơ quan, tổ chức, đơn vị;

c) Thông báo bằng văn bản đến cơ quan, tổ chức, đơn vị, cá nhân có liên quan;

d) Phát hành ấn phẩm;

đ) Thông báo trên các phương tiện thông tin đại chúng;

e) Đưa lên trang thông tin điện tử;

g) Cung cấp thông tin theo yêu cầu của cơ quan, tổ chức, cá nhân.

2. Ngoài những trường hợp pháp luật có quy định về hình thức công khai, người đứng đầu cơ quan, tổ chức, đơn vị có trách nhiệm lựa chọn một hoặc một số hình thức công khai quy định tại khoản 1 Điều này.

Điều 13. Công khai, minh bạch trong mua sắm công và xây dựng cơ bản

1. Việc mua sắm công và xây dựng cơ bản phải được công khai theo quy định của pháp luật.

2. Trường hợp mua sắm công và xây dựng cơ bản mà pháp luật quy định phải đấu thầu thì nội dung công khai bao gồm:

a) Kế hoạch đấu thầu, mời sơ tuyển và kết quả sơ tuyển, mời thầu;

b) Danh mục nhà thầu tham gia đấu thầu hạn chế, danh sách ngắn nhà thầu tham gia đấu thầu hạn chế, kết quả lựa chọn nhà thầu;

c) Thông tin về cá nhân, tổ chức thuộc chủ dự án, bên mời thầu, nhà thầu, cơ quan quản lý hoặc đối tượng khác vi phạm pháp luật về đấu thầu; thông tin về nhà thầu bị cấm tham gia và thông tin về xử lý vi phạm pháp luật về đấu thầu;

d) Văn bản quy phạm pháp luật về đấu thầu, hệ thống thông tin dữ liệu về đấu thầu;

đ) Báo cáo tổng kết công tác đấu thầu trên phạm vi toàn quốc của Bộ Kế hoạch và Đầu tư; báo cáo tổng kết công tác đấu thầu của bộ, ngành, địa phương và cơ sở;

e) Thẩm quyền, thủ tục tiếp nhận và giải quyết khiếu nại, tố cáo trong đấu thầu.

Điều 14. Công khai, minh bạch trong quản lý dự án đầu tư xây dựng

1. Dự án quy hoạch đầu tư xây dựng phải được lấy ý kiến của nhân dân địa phương nơi quy hoạch.

2. Dự án đầu tư xây dựng từ ngân sách địa phương phải được Hội đồng nhân dân xem xét, quyết định.

3. Dự án đầu tư xây dựng sau khi được quyết định, phê duyệt phải được công khai để nhân dân giám sát.

Điều 15. Công khai, minh bạch về tài chính và ngân sách nhà nước

1. Các cấp ngân sách, đơn vị dự toán ngân sách phải công khai chi tiết số liệu dự toán và quyết toán đã được cơ quan nhà nước có thẩm quyền quyết định, phê chuẩn, kể cả khoản ngân sách bổ sung.

2. Đơn vị dự toán ngân sách có nguồn thu và các khoản chi từ các khoản đóng góp của tổ chức, cá nhân theo quy định của pháp luật phải công khai mục đích huy động, kết quả huy động và hiệu quả việc sử dụng các nguồn huy động.

3. Tổ chức được ngân sách nhà nước hỗ trợ phải công khai các nội dung sau đây:

a) Số liệu dự toán, quyết toán;

b) Khoản đóng góp của tổ chức, cá nhân (nếu có);

c) Cơ sở xác định mức hỗ trợ và số tiền ngân sách nhà nước hỗ trợ.

4. Dự án đầu tư xây dựng cơ bản có sử dụng vốn từ ngân sách nhà nước phải công khai các nội dung sau đây:

a) Việc phân bổ vốn đầu tư trong dự toán ngân sách nhà nước được giao hằng năm cho các dự án;

b) Dự toán ngân sách của dự án đầu tư theo kế hoạch đầu tư được duyệt, mức vốn đầu tư của dự án được giao trong dự toán ngân sách năm;

c) Quyết toán vốn đầu tư của dự án hằng năm;

d) Quyết toán vốn đầu tư khi dự án hoàn thành đã được cấp có thẩm quyền phê duyệt.

5. Quỹ có nguồn từ ngân sách nhà nước phải công khai các nội dung sau đây:

a) Quy chế hoạt động và cơ chế tài chính của quỹ;

b) Kế hoạch tài chính hằng năm, trong đó chi tiết các khoản thu, chi có quan hệ với ngân sách nhà nước theo quy định của cấp có thẩm quyền;

c) Kết quả hoạt động của quỹ;

d) Quyết toán năm được cấp có thẩm quyền phê duyệt.

6. Việc phân bổ, sử dụng ngân sách, tài sản của Nhà nước cho các dự án, chương trình mục tiêu đã được cơ quan nhà nước có thẩm quyền phê duyệt phải công khai cho cơ quan, tổ chức, đơn vị hữu quan và nhân dân nơi trực tiếp thụ hưởng biết.

Điều 16. Công khai, minh bạch việc huy động và sử dụng các khoản đóng góp của nhân dân

1. Việc huy động các khoản đóng góp của nhân dân để đầu tư xây dựng công trình, lập quỹ trong phạm vi địa phương phải lấy ý kiến nhân dân và được Hội đồng nhân dân cùng cấp xem xét, quyết định.

2. Việc huy động, sử dụng các khoản đóng góp của nhân dân quy định tại khoản 1 Điều này phải được công khai để nhân dân giám sát và phải chịu sự thanh tra, kiểm tra, giám sát theo quy định của pháp luật.

3. Nội dung phải công khai bao gồm mục đích huy động, mức đóng góp, việc sử dụng, kết quả sử dụng và báo cáo quyết toán.

4. Công trình cơ sở hạ tầng tại xã, phường, thị trấn sử dụng các khoản đóng góp của nhân dân phải công khai các nội dung sau đây:

a) Nội dung phải công khai quy định tại khoản 3 Điều này;

b) Dự toán cho từng công trình theo kế hoạch đầu tư được duyệt;

c) Nguồn vốn đầu tư cho từng công trình;

d) Kết quả đã huy động của từng đối tượng cụ thể, thời gian huy động;

đ) Kết quả lựa chọn nhà thầu đã được cấp có thẩm quyền phê duyệt;

e) Tiến độ thi công và kết quả nghiệm thu khối lượng, chất lượng công trình và quyết toán công trình.

5. Việc huy động, sử dụng các khoản đóng góp của nhân dân vì mục đích từ thiện, nhân đạo được thực hiện theo quy định tại khoản 2 và khoản 3 Điều này.

Điều 17. Công khai, minh bạch việc quản lý, sử dụng các khoản hỗ trợ, viện trợ

Việc quản lý, phân bổ, sử dụng nguồn vốn hỗ trợ phát triển chính thức (ODA) được thực hiện theo quy định tại Điều 15 của Luật này. Đối với các khoản viện trợ phi chính phủ phải được công khai cho các đối tượng thụ hưởng biết.

Điều 18. Công khai, minh bạch trong quản lý doanh nghiệp của Nhà nước

Doanh nghiệp của Nhà nước có trách nhiệm công khai vốn và tài sản của Nhà nước đầu tư vào doanh nghiệp, vốn vay ưu đãi, báo cáo tài chính và kết quả kiểm toán, việc trích, lập và sử dụng quỹ của doanh nghiệp, việc tuyển dụng lao động, bổ nhiệm các chức danh lãnh đạo, quản lý của doanh nghiệp và nội dung khác theo quy định của pháp luật.

Điều 19. Công khai, minh bạch trong cổ phần hoá doanh nghiệp của Nhà nước

1. Việc cổ phần hoá doanh nghiệp của Nhà nước phải công khai, minh bạch; không được cổ phần hoá khép kín trong nội bộ doanh nghiệp. Doanh nghiệp được cổ phần hoá có trách nhiệm công khai thực trạng tài chính khi xác định giá trị doanh nghiệp.

2. Cơ quan nhà nước có thẩm quyền có trách nhiệm công khai giá trị doanh nghiệp được cổ phần hoá và việc điều chỉnh giá trị doanh nghiệp (nếu có).

3. Việc bán cổ phần lần đầu của doanh nghiệp được cổ phần hoá phải thực hiện bằng phương thức bán đấu giá.

Điều 20. Kiểm toán việc sử dụng ngân sách, tài sản của Nhà nước

1. Cơ quan, tổ chức, đơn vị có trách nhiệm thực hiện kiểm toán và chịu sự kiểm toán việc sử dụng ngân sách, tài sản của Nhà nước theo quy định của pháp luật về kiểm toán.

2. Báo cáo kiểm toán phải được công khai theo quy định tại Điều 12 của Luật này.

Điều 21. Công khai, minh bạch trong quản lý và sử dụng đất

1. Việc lập quy hoạch, kế hoạch sử dụng đất phải bảo đảm dân chủ và công khai.

2. Trong quá trình lập và điều chỉnh quy hoạch, kế hoạch sử dụng đất chi tiết, cơ quan, tổ chức thực hiện việc lập quy hoạch, kế hoạch đó phải thông báo công khai cho nhân dân địa phương nơi được quy hoạch, điều chỉnh biết.

3. Quy hoạch, kế hoạch sử dụng đất chi tiết, việc giải phóng mặt bằng, giá đền bù khi thu hồi đất sau khi được cơ quan nhà nước có thẩm quyền quyết định, phê duyệt hoặc điều chỉnh phải được công khai.

4. Thẩm quyền, trình tự, thủ tục và việc cấp giấy chứng nhận quyền sử dụng đất; quy hoạch chi tiết và việc phân lô đất ở, đối tượng được giao đất làm nhà ở phải được công khai.

Điều 22. Công khai, minh bạch trong quản lý, sử dụng nhà ở

1. Thẩm quyền, trình tự, thủ tục và việc cấp giấy phép xây dựng nhà ở và giấy chứng nhận quyền sở hữu nhà ở phải được công khai.

2. Việc hoá giá nhà ở thuộc sở hữu nhà nước, đối tượng được hoá giá nhà ở và các khoản tiền phải nộp khi hoá giá nhà ở phải được công khai.

3. Việc bán nhà ở cho người tái định cư, người có thu nhập thấp và những đối tượng ưu tiên khác phải được công khai.

Điều 23. Công khai, minh bạch trong lĩnh vực giáo dục

1. Việc tuyển sinh, thi, kiểm tra, cấp văn bằng, chứng chỉ phải được công khai.

2. Cơ quan quản lý giáo dục, cơ sở giáo dục có sử dụng ngân sách, tài sản của Nhà nước phải công khai việc quản lý, sử dụng ngân sách, tài sản của Nhà nước, việc thu, quản lý, sử dụng học phí, lệ phí tuyển sinh, các khoản thu từ hoạt động tư vấn, chuyển giao công nghệ, các khoản hỗ trợ, đầu tư cho giáo dục và các khoản thu khác theo quy định của pháp luật.

Điều 24. Công khai, minh bạch trong lĩnh vực y tế

1. Thẩm quyền, trình tự, thủ tục và việc cấp, thu hồi chứng chỉ hành nghề y, dược tư nhân, giấy chứng nhận đủ điều kiện hành nghề cho các cơ sở hành nghề y, dược phải được công khai.

2. Cơ quan quản lý y tế, cơ sở khám, chữa bệnh có sử dụng ngân sách, tài sản của Nhà nước phải công khai việc thu, quản lý, sử dụng ngân sách, tài sản của Nhà nước, giá thuốc, việc thu, quản lý, sử dụng các loại phí liên quan đến việc khám, chữa bệnh và các khoản thu khác theo quy định của pháp luật.

Điều 25. Công khai, minh bạch trong lĩnh vực khoa học - công nghệ

1. Việc xét, tuyển chọn, giao trực tiếp, tài trợ thực hiện nhiệm vụ khoa học - công nghệ và việc đánh giá, nghiệm thu kết quả thực hiện nhiệm vụ khoa học - công nghệ phải được tiến hành công khai.

2. Cơ quan quản lý khoa học - công nghệ, đơn vị nghiên cứu khoa học - công nghệ phải công khai việc quản lý, sử dụng ngân sách, tài sản của Nhà nước, các khoản hỗ trợ, viện trợ, đầu tư, các khoản thu từ hoạt động khoa học - công nghệ.

Điều 26. Công khai, minh bạch trong lĩnh vực thể dục, thể thao

Cơ quan quản lý thể dục, thể thao, Uỷ ban Ô-lim-pích Việt Nam, các liên đoàn thể thao, cơ sở thể dục, thể thao có trách nhiệm công khai việc quản lý, sử dụng ngân sách, tài sản của Nhà nước, các khoản thu từ hoạt động và dịch vụ thể dục, thể thao, khoản tài trợ, hỗ trợ, đóng góp của tổ chức, cá nhân trong nước và nước ngoài cho hoạt động thể dục, thể thao.

Điều 27. Công khai, minh bạch trong hoạt động thanh tra, giải quyết khiếu nại, tố cáo, kiểm toán nhà nước

1. Hoạt động thanh tra, giải quyết khiếu nại, tố cáo, kiểm toán nhà nước phải được tiến hành công khai theo quy định của pháp luật.

2. Văn bản, quyết định sau đây phải được công khai, trừ trường hợp pháp luật có quy định khác:

a) Kết luận thanh tra;

b) Quyết định giải quyết khiếu nại, quyết định giải quyết tố cáo;

c) Báo cáo kiểm toán.

Điều 28. Công khai, minh bạch trong hoạt động giải quyết các công việc của cơ quan, tổ chức, đơn vị, cá nhân

1. Cơ quan, tổ chức, cá nhân có thẩm quyền quản lý trong lĩnh vực nhà, đất, xây dựng, đăng ký kinh doanh, xét duyệt dự án, cấp vốn ngân sách nhà nước, tín dụng, ngân hàng, xuất khẩu, nhập khẩu, xuất cảnh, nhập cảnh, quản lý hộ khẩu, thuế, hải quan, bảo hiểm và các cơ quan, tổ chức, cá nhân khác trực tiếp giải quyết công việc của cơ quan, tổ chức, đơn vị, cá nhân phải công khai thủ tục hành chính, giải quyết đúng thời hạn, đúng pháp luật và đúng yêu cầu hợp pháp của cơ quan, tổ chức, đơn vị, cá nhân.

2. Cơ quan, tổ chức, đơn vị, cá nhân có quyền đề nghị với cơ quan, tổ chức, cá nhân có thẩm quyền giải quyết công việc của mình giải thích rõ những nội dung có liên quan. Khi nhận được đề nghị của cơ quan, tổ chức, đơn vị, cá nhân thì cơ quan, tổ chức, cá nhân có thẩm quyền phải kịp thời giải thích công khai.

3. Trong trường hợp cơ quan, tổ chức, cá nhân có thẩm quyền giải thích chưa thoả đáng hoặc cố tình gây khó khăn, phiền hà thì cơ quan, tổ chức, đơn vị, cá nhân có quyền kiến nghị lên cơ quan, tổ chức cấp trên trực tiếp của cơ quan, tổ chức, cá nhân đó.

Điều 29. Công khai, minh bạch trong lĩnh vực tư pháp

Việc thụ lý, điều tra, truy tố, kiểm sát, xét xử, thi hành án phải được công khai theo quy định của pháp luật về tố tụng và các quy định khác của pháp luật có liên quan.

Điều 30. Công khai, minh bạch trong công tác tổ chức - cán bộ

1. Việc tuyển dụng cán bộ, công chức, viên chức và người lao động khác vào cơ quan, tổ chức, đơn vị phải được công khai về số lượng, tiêu chuẩn, hình thức và kết quả tuyển dụng.

2. Việc quy hoạch, đào tạo, bổ nhiệm, chuyển ngạch, luân chuyển, điều động, khen thưởng, cho thôi việc, cho thôi giữ chức vụ, miễn nhiệm, bãi nhiệm, kỷ luật, hưu trí đối với cán bộ, công chức, viên chức và người lao động khác phải được công khai trong cơ quan, tổ chức, đơn vị nơi người đó làm việc.

Điều 31. Quyền yêu cầu cung cấp thông tin của cơ quan, tổ chức

1. Cơ quan nhà nước, tổ chức chính trị, tổ chức chính trị - xã hội, cơ quan báo chí trong phạm vi nhiệm vụ, quyền hạn của mình có quyền yêu cầu cơ quan, tổ chức, đơn vị có trách nhiệm cung cấp thông tin về hoạt động của cơ quan, tổ chức, đơn vị mình theo quy định của pháp luật.

2. Trong thời hạn mười ngày, kể từ ngày nhận được yêu cầu, cơ quan, tổ chức, đơn vị được yêu cầu phải cung cấp thông tin, trừ trường hợp nội dung thông tin đã được công khai trên các phương tiện thông tin đại chúng, được phát hành ấn phẩm hoặc niêm yết công khai; trường hợp không cung cấp hoặc chưa cung cấp được thì phải trả lời bằng văn bản cho cơ quan, tổ chức yêu cầu biết và nêu rõ lý do.

Điều 32. Quyền yêu cầu cung cấp thông tin của cá nhân

1. Cán bộ, công chức, viên chức và người lao động khác có quyền yêu cầu người đứng đầu cơ quan, tổ chức, đơn vị nơi mình làm việc cung cấp thông tin về hoạt động của cơ quan, tổ chức, đơn vị đó.

2. Công dân có quyền yêu cầu Chủ tịch Uỷ ban nhân dân xã, phường, thị trấn nơi mình cư trú cung cấp thông tin về hoạt động của Uỷ ban nhân dân xã, phường, thị trấn đó.

3. Trong thời hạn mười ngày, kể từ ngày nhận được yêu cầu, người được yêu cầu có trách nhiệm cung cấp thông tin, trừ trường hợp nội dung thông tin đã được công khai trên các phương tiện thông tin đại chúng, được phát hành ấn phẩm hoặc niêm yết công khai; trường hợp không cung cấp hoặc chưa cung cấp được thì phải trả lời bằng văn bản cho người yêu cầu biết và nêu rõ lý do.

Điều 33. Công khai báo cáo hằng năm về phòng, chống tham nhũng

1. Hằng năm, Chính phủ có trách nhiệm báo cáo Quốc hội về công tác phòng, chống tham nhũng trong phạm vi cả nước; Uỷ ban nhân dân có trách nhiệm báo cáo Hội đồng nhân dân cùng cấp về công tác phòng, chống tham nhũng ở địa phương.

2. Báo cáo về công tác phòng, chống tham nhũng phải được công khai.

Mục 2

XÂY DỰNG VÀ THỰC HIỆN CÁC CHẾ ĐỘ, ĐỊNH MỨC, TIÊU CHUẨN

Điều 34. Xây dựng, ban hành và thực hiện các chế độ, định mức, tiêu chuẩn

1. Cơ quan nhà nước trong phạm vi nhiệm vụ, quyền hạn của mình có trách nhiệm:

a) Xây dựng, ban hành và công khai các chế độ, định mức, tiêu chuẩn;

b) Công khai các quy định về chế độ, định mức, tiêu chuẩn về quyền lợi đối với từng loại chức danh trong cơ quan mình;

c) Chấp hành nghiêm chỉnh các quy định về chế độ, định mức, tiêu chuẩn.

2. Tổ chức chính trị, tổ chức chính trị - xã hội, đơn vị sự nghiệp và các cơ quan, tổ chức, đơn vị khác có sử dụng ngân sách nhà nước căn cứ vào quy định tại khoản 1 Điều này hướng dẫn áp dụng hoặc phối hợp với cơ quan nhà nước có thẩm quyền xây dựng, ban hành và công khai các chế độ, định mức, tiêu chuẩn áp dụng trong cơ quan, tổ chức, đơn vị mình.

3. Nghiêm cấm cơ quan, tổ chức, đơn vị ban hành trái pháp luật các chế độ, định mức, tiêu chuẩn.

Điều 35. Kiểm tra và xử lý vi phạm quy định về chế độ, định mức, tiêu chuẩn

1. Cơ quan, tổ chức, đơn vị phải thường xuyên kiểm tra việc chấp hành và xử lý kịp thời hành vi vi phạm quy định về chế độ, định mức, tiêu chuẩn.

2. Người có hành vi vi phạm quy định về chế độ, định mức, tiêu chuẩn phải bị xử lý theo quy định của pháp luật.

3. Người cho phép sử dụng vượt chế độ, định mức, tiêu chuẩn phải bồi thường phần giá trị mà mình cho phép sử dụng vượt quá; người sử dụng vượt chế độ, định mức, tiêu chuẩn có trách nhiệm liên đới bồi thường phần giá trị được sử dụng vượt quá.

4. Người cho phép thực hiện chế độ, định mức, tiêu chuẩn chuyên môn - kỹ thuật thấp hơn mức quy định phải bồi thường phần giá trị mà mình cho phép sử dụng thấp hơn; người hưởng lợi từ việc thực hiện chế độ, định mức, tiêu chuẩn chuyên môn - kỹ thuật thấp hơn có trách nhiệm liên đới bồi thường phần giá trị được hưởng lợi.

Mục 3

QUY TẮC ỨNG XỬ, QUY TẮC ĐẠO ĐỨC NGHỀ NGHIỆP

VIỆC CHUYỂN ĐỔI VỊ TRÍ CÔNG TÁC CỦA CÁN BỘ, CÔNG CHỨC, VIÊN CHỨC

Điều 36. Quy tắc ứng xử của cán bộ, công chức, viên chức

1. Quy tắc ứng xử là các chuẩn mực xử sự của cán bộ, công chức, viên chức trong thi hành nhiệm vụ, công vụ và trong quan hệ xã hội, bao gồm những việc phải làm hoặc không được làm, phù hợp với đặc thù công việc của từng nhóm cán bộ, công chức, viên chức và từng lĩnh vực hoạt động công vụ, nhằm bảo đảm sự liêm chính và trách nhiệm của cán bộ, công chức, viên chức.

2. Quy tắc ứng xử của cán bộ, công chức, viên chức được công khai để nhân dân giám sát việc chấp hành.

Điều 37. Những việc cán bộ, công chức, viên chức không được làm

1. Cán bộ, công chức, viên chức không được làm những việc sau đây:

a) Cửa quyền, hách dịch, gây khó khăn, phiền hà đối với cơ quan, tổ chức, đơn vị, cá nhân trong khi giải quyết công việc;

b) Thành lập, tham gia thành lập hoặc tham gia quản lý, điều hành doanh nghiệp tư nhân, công ty trách nhiệm hữu hạn, công ty cổ phần, công ty hợp danh, hợp tác xã, bệnh viện tư, trường học tư và tổ chức nghiên cứu khoa học tư, trừ trường hợp pháp luật có quy định khác;

c) Làm tư vấn cho doanh nghiệp, tổ chức, cá nhân khác ở trong nước và nước ngoài về các công việc có liên quan đến bí mật nhà nước, bí mật công tác, những công việc thuộc thẩm quyền giải quyết của mình hoặc mình tham gia giải quyết;

d) Kinh doanh trong lĩnh vực mà trước đây mình có trách nhiệm quản lý sau khi thôi giữ chức vụ trong một thời hạn nhất định theo quy định của Chính phủ;

đ) Sử dụng trái phép thông tin, tài liệu của cơ quan, tổ chức, đơn vị vì vụ lợi.

2. Người đứng đầu, cấp phó của người đứng đầu cơ quan, vợ hoặc chồng của những người đó không được góp vốn vào doanh nghiệp hoạt động trong phạm vi ngành, nghề mà người đó trực tiếp thực hiện việc quản lý nhà nước.

3. Người đứng đầu, cấp phó của người đứng đầu cơ quan, tổ chức, đơn vị không được bố trí vợ hoặc chồng, bố, mẹ, con, anh, chị, em ruột của mình giữ chức vụ quản lý về tổ chức nhân sự, kế toán - tài vụ, làm thủ quỹ, thủ kho trong cơ quan, tổ chức, đơn vị hoặc giao dịch, mua bán vật tư, hàng hoá, ký kết hợp đồng cho cơ quan, tổ chức, đơn vị đó.

4. Người đứng đầu, cấp phó của người đứng đầu cơ quan không được để vợ hoặc chồng, bố, mẹ, con kinh doanh trong phạm vi do mình quản lý trực tiếp.

5. Cán bộ, công chức, viên chức là thành viên Hội đồng quản trị, Tổng giám đốc, Phó tổng giám đốc, Giám đốc, Phó giám đốc, Kế toán trưởng và những cán bộ quản lý khác trong doanh nghiệp của Nhà nước không được ký kết hợp đồng với doanh nghiệp thuộc sở hữu của vợ hoặc chồng, bố, mẹ, con, anh, chị, em ruột; cho phép doanh nghiệp thuộc sở hữu của vợ hoặc chồng, bố, mẹ, con, anh, chị, em ruột tham dự các gói thầu của doanh nghiệp mình; bố trí vợ hoặc chồng, bố, mẹ, con, anh, chị, em ruột giữ chức vụ quản lý về tổ chức nhân sự, kế toán - tài vụ, làm thủ quỹ, thủ kho trong doanh nghiệp hoặc giao dịch, mua bán vật tư, hàng hoá, ký kết hợp đồng cho doanh nghiệp.

6. Quy định tại các khoản 1, 2, 3 và 4 Điều này cũng được áp dụng đối với các đối tượng sau đây:

a) Sĩ quan, quân nhân chuyên nghiệp, công nhân quốc phòng trong cơ quan, đơn vị thuộc Quân đội nhân dân;

b) Sĩ quan, hạ sĩ quan nghiệp vụ, sĩ quan, hạ sĩ quan chuyên môn - kỹ thuật trong cơ quan, đơn vị thuộc Công an nhân dân.

Điều 38. Nghĩa vụ báo cáo và xử lý báo cáo về dấu hiệu tham nhũng

1. Khi phát hiện có dấu hiệu tham nhũng trong cơ quan, tổ chức, đơn vị nơi mình làm việc thì cán bộ, công chức, viên chức phải báo cáo ngay với người đứng đầu cơ quan, tổ chức, đơn vị đó; trường hợp người đứng đầu cơ quan, tổ chức, đơn vị có liên quan đến dấu hiệu tham nhũng đó thì báo cáo với người đứng đầu cơ quan, tổ chức, đơn vị cấp trên trực tiếp.

2. Chậm nhất là mười ngày, kể từ ngày nhận được báo cáo về dấu hiệu tham nhũng, người được báo cáo phải xử lý vụ việc theo thẩm quyền hoặc chuyển cho cơ quan, tổ chức, cá nhân có thẩm quyền xem xét xử lý và thông báo cho người báo cáo; đối với vụ việc phức tạp thì thời hạn trên có thể kéo dài nhưng không quá ba mươi ngày; trường hợp cần thiết thì quyết định hoặc đề nghị người có thẩm quyền quyết định áp dụng các biện pháp nhằm ngăn chặn, khắc phục hậu quả của hành vi tham nhũng và bảo vệ người báo cáo.

Điều 39. Trách nhiệm của người không báo cáo hoặc không xử lý báo cáo về dấu hiệu tham nhũng

Cán bộ, công chức, viên chức biết được hành vi tham nhũng mà không báo cáo, người nhận được báo cáo về dấu hiệu tham nhũng mà không xử lý thì phải chịu trách nhiệm theo quy định của pháp luật.

Điều 40. Việc tặng quà và nhận quà tặng của cán bộ, công chức, viên chức

1. Cơ quan, tổ chức, đơn vị không được sử dụng ngân sách, tài sản của Nhà nước làm quà tặng, trừ trường hợp pháp luật có quy định khác.

2. Cán bộ, công chức, viên chức không được nhận tiền, tài sản hoặc lợi ích vật chất khác của cơ quan, tổ chức, đơn vị, cá nhân liên quan đến công việc do mình giải quyết hoặc thuộc phạm vi quản lý của mình.

3. Nghiêm cấm lợi dụng việc tặng quà, nhận quà tặng để hối lộ hoặc thực hiện các hành vi khác vì vụ lợi.

4. Chính phủ quy định chi tiết việc tặng quà, nhận quà tặng và nộp lại quà tặng của cán bộ, công chức, viên chức.

Điều 41. Thẩm quyền ban hành quy tắc ứng xử của cán bộ, công chức, viên chức

1. Bộ trưởng, Thủ trưởng cơ quan ngang bộ, Thủ trưởng cơ quan thuộc Chính phủ, Chủ nhiệm Văn phòng Quốc hội, Chủ nhiệm Văn phòng Chủ tịch nước ban hành quy tắc ứng xử của cán bộ, công chức, viên chức làm việc trong cơ quan, ngành, lĩnh vực do mình quản lý.

2. Chánh án Toà án nhân dân tối cao, Viện trưởng Viện kiểm sát nhân dân tối cao ban hành quy tắc ứng xử của Thẩm phán, Hội thẩm, Thư ký Toà án, Kiểm sát viên và cán bộ, công chức, viên chức khác trong cơ quan Toà án, Viện kiểm sát.

3. Bộ trưởng Bộ Nội vụ ban hành quy tắc ứng xử của cán bộ, công chức, viên chức làm việc trong bộ máy chính quyền địa phương; phối hợp với cơ quan trung ương của tổ chức chính trị - xã hội ban hành quy tắc ứng xử của cán bộ, công chức, viên chức trong tổ chức này.

Điều 42. Quy tắc đạo đức nghề nghiệp

1. Quy tắc đạo đức nghề nghiệp là chuẩn mực xử sự phù hợp với đặc thù của từng nghề bảo đảm sự liêm chính, trung thực và trách nhiệm trong việc hành nghề.

2. Tổ chức xã hội - nghề nghiệp phối hợp với cơ quan nhà nước có thẩm quyền ban hành quy tắc đạo đức nghề nghiệp đối với hội viên của mình theo quy định của pháp luật.

Điều 43. Chuyển đổi vị trí công tác của cán bộ, công chức, viên chức

1. Cơ quan, tổ chức, đơn vị theo thẩm quyền quản lý có trách nhiệm thực hiện việc định kỳ chuyển đổi cán bộ, công chức, viên chức làm việc tại một số vị trí liên quan đến việc quản lý ngân sách, tài sản của Nhà nước, trực tiếp tiếp xúc và giải quyết công việc của cơ quan, tổ chức, đơn vị, cá nhân nhằm chủ động phòng ngừa tham nhũng.

2. Việc chuyển đổi vị trí công tác phải theo kế hoạch và được công khai trong nội bộ cơ quan, tổ chức, đơn vị.

3. Việc chuyển đổi vị trí công tác quy định tại khoản 1 và khoản 2 Điều này chỉ áp dụng đối với cán bộ, công chức, viên chức không giữ chức vụ quản lý. Việc luân chuyển cán bộ, công chức giữ chức vụ quản lý thực hiện theo quy định về luân chuyển cán bộ.

4. Chính phủ ban hành Danh mục các vị trí công tác và thời hạn định kỳ chuyển đổi quy định tại khoản 1 Điều này.

Mục 4

MINH BẠCH TÀI SẢN, THU NHẬP

Điều 44. Nghĩa vụ kê khai tài sản

1. Những người sau đây phải kê khai tài sản:

a) Cán bộ từ Phó trưởng phòng của Uỷ ban nhân dân huyện, quận, thị xã, thành phố thuộc tỉnh trở lên và tương đương trong các cơ quan, tổ chức, đơn vị;

b) Một số cán bộ, công chức tại xã, phường, thị trấn; người làm công tác quản lý ngân sách, tài sản của Nhà nước hoặc trực tiếp tiếp xúc và giải quyết công việc của cơ quan, tổ chức, đơn vị, cá nhân;

c) Người ứng cử đại biểu Quốc hội, đại biểu Hội đồng nhân dân.

Chính phủ quy định cụ thể những người phải kê khai tài sản quy định tại khoản này.

2. Người có nghĩa vụ kê khai tài sản phải kê khai tài sản, mọi biến động về tài sản thuộc sở hữu của mình và tài sản thuộc sở hữu của vợ hoặc chồng và con chưa thành niên.

3. Người có nghĩa vụ kê khai tài sản phải kê khai trung thực và chịu trách nhiệm về việc kê khai.

Điều 45. Tài sản phải kê khai

Các loại tài sản phải kê khai bao gồm:

1. Nhà, quyền sử dụng đất;

2. Kim khí quý, đá quý, tiền, giấy tờ có giá và các loại tài sản khác mà giá trị của mỗi loại từ năm mươi triệu đồng trở lên;

3. Tài sản, tài khoản ở nước ngoài;

4. Thu nhập phải chịu thuế theo quy định của pháp luật.

Điều 46. Thủ tục kê khai tài sản

1. Việc kê khai tài sản được thực hiện hằng năm tại cơ quan, tổ chức, đơn vị nơi người có nghĩa vụ kê khai làm việc và được hoàn thành chậm nhất vào ngày 31 tháng 12.

2. Người có nghĩa vụ kê khai tài sản phải ghi rõ những thay đổi về tài sản so với lần kê khai trước đó.

3. Bản kê khai tài sản được nộp cho cơ quan, tổ chức, đơn vị có thẩm quyền quản lý người có nghĩa vụ kê khai tài sản.

Điều 47. Xác minh tài sản

1. Việc xác minh tài sản chỉ được thực hiện khi có quyết định của cơ quan, tổ chức có thẩm quyền quản lý người có nghĩa vụ kê khai tài sản.

2. Việc xác minh tài sản được thực hiện trong các trường hợp sau đây:

a) Phục vụ cho việc bầu cử, bổ nhiệm, cách chức, miễn nhiệm, bãi nhiệm hoặc kỷ luật đối với người có nghĩa vụ kê khai tài sản khi xét thấy cần thiết;

b) Theo yêu cầu của Hội đồng bầu cử hoặc cơ quan, tổ chức có thẩm quyền;

c) Có hành vi tham nhũng.

Điều 48. Thủ tục xác minh tài sản

1. Trước khi ra quyết định xác minh tài sản, cơ quan, tổ chức có thẩm quyền yêu cầu người có nghĩa vụ kê khai giải trình rõ việc kê khai. Việc giải trình phải được thực hiện trong thời hạn năm ngày, kể từ ngày nhận được yêu cầu giải trình.

2. Cơ quan, tổ chức có thẩm quyền ra quyết định xác minh trong thời hạn năm ngày, kể từ ngày phát sinh căn cứ quy định tại khoản 2 Điều 47 của Luật này.

3. Cơ quan, tổ chức, đơn vị, cá nhân hữu quan có trách nhiệm cung cấp thông tin, tài liệu phục vụ cho việc xác minh khi có yêu cầu của cơ quan, tổ chức có thẩm quyền.

4. Trong thời hạn hai mươi ngày kể từ ngày ra quyết định xác minh, cơ quan, tổ chức, đơn vị quản lý người có nghĩa vụ kê khai tài sản tiến hành thẩm tra, xác minh và phải ra kết luận về sự minh bạch trong kê khai tài sản.

5. Thủ tục xác minh tài sản của người có tên trong danh sách ứng cử đại biểu Quốc hội, đại biểu Hội đồng nhân dân được thực hiện theo quy định tại các khoản 1, 2, 3 và 4 Điều này. Thời hạn xác minh phải đáp ứng yêu cầu về thời gian bầu cử đại biểu Quốc hội, đại biểu Hội đồng nhân dân.

Điều 49. Kết luận về sự minh bạch trong kê khai tài sản

1. Kết luận về sự minh bạch trong kê khai tài sản là kết luận về tính trung thực của việc kê khai tài sản.

2. Kết luận về sự minh bạch trong kê khai tài sản phải được gửi cho cơ quan, tổ chức yêu cầu xác minh và người có tài sản được xác minh.

3. Cơ quan, tổ chức, đơn vị quy định tại khoản 4 Điều 48 của Luật này phải chịu trách nhiệm về tính khách quan, chính xác và nội dung kết luận của mình.

Điều 50. Công khai kết luận về sự minh bạch trong kê khai tài sản

1. Khi có yêu cầu và theo quyết định của cơ quan, tổ chức có thẩm quyền, bản kết luận về sự minh bạch trong kê khai tài sản được công khai tại các địa điểm sau đây:

a) Trong cơ quan, tổ chức, đơn vị nơi người có nghĩa vụ kê khai tài sản làm việc khi người đó được bổ nhiệm, bầu, phê chuẩn;

b) Tại hội nghị cử tri nơi công tác, nơi cư trú đối với người ứng cử đại biểu Quốc hội, đại biểu Hội đồng nhân dân;

c) Tại cơ quan, tổ chức, đơn vị nơi người được đề nghị để Quốc hội, Hội đồng nhân dân hoặc Đại hội của tổ chức chính trị, tổ chức chính trị - xã hội bầu, phê chuẩn.

2. Kết luận về sự minh bạch trong kê khai tài sản của người bị khởi tố về hành vi tham nhũng phải được công khai trong cơ quan, tổ chức, đơn vị nơi người đó làm việc.

Điều 51. Trách nhiệm của cơ quan, tổ chức, đơn vị quản lý người có nghĩa vụ kê khai tài sản

Cơ quan, tổ chức, đơn vị có trách nhiệm quản lý và lưu giữ bản kê khai tài sản của người có nghĩa vụ kê khai do mình quản lý; tổ chức việc xác minh theo quyết định của cơ quan, tổ chức có thẩm quyền; kết luận về sự minh bạch trong kê khai tài sản và công khai kết luận đó theo quyết định của cơ quan, tổ chức có thẩm quyền trong các trường hợp quy định tại Điều 50 của Luật này.

Điều 52. Xử lý người kê khai tài sản không trung thực

1. Người kê khai tài sản không trung thực bị xử lý kỷ luật theo quy định của pháp luật. Quyết định kỷ luật đối với người kê khai tài sản không trung thực phải được công khai tại cơ quan, tổ chức, đơn vị nơi người đó làm việc.

2. Người ứng cử đại biểu Quốc hội, đại biểu Hội đồng nhân dân mà kê khai tài sản không trung thực thì bị xoá tên khỏi danh sách những người ứng cử; người được dự kiến bổ nhiệm, phê chuẩn mà kê khai tài sản không trung thực thì không được bổ nhiệm, phê chuẩn vào chức vụ đã dự kiến.

Điều 53. Kiểm soát thu nhập

Chính phủ trình Quốc hội ban hành văn bản quy phạm pháp luật về kiểm soát thu nhập của người có chức vụ, quyền hạn.

Mục 5

CHẾ ĐỘ TRÁCH NHIỆM CỦA NGƯỜI ĐỨNG ĐẦU CƠ QUAN,

TỔ CHỨC, ĐƠN VỊ KHI ĐỂ XẢY RA THAM NHŨNG

Điều 54. Trách nhiệm của người đứng đầu cơ quan, tổ chức, đơn vị khi để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách

1. Người đứng đầu cơ quan, tổ chức, đơn vị phải chịu trách nhiệm về việc để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách.

Người đứng đầu cơ quan, tổ chức, đơn vị phải chịu trách nhiệm trực tiếp về việc để xảy ra hành vi tham nhũng của người do mình trực tiếp quản lý, giao nhiệm vụ.

2. Cấp phó của người đứng đầu cơ quan, tổ chức, đơn vị phải chịu trách nhiệm trực tiếp về việc để xảy ra hành vi tham nhũng trong lĩnh vực công tác và trong đơn vị do mình trực tiếp phụ trách.

Người đứng đầu cơ quan, tổ chức, đơn vị phải chịu trách nhiệm liên đới về việc để xảy ra hành vi tham nhũng trong lĩnh vực công tác và trong đơn vị do cấp phó của mình trực tiếp phụ trách.

3. Người đứng đầu đơn vị trực thuộc cơ quan, tổ chức phải chịu trách nhiệm trực tiếp về việc để xảy ra hành vi tham nhũng trong đơn vị do mình quản lý.

4. Việc xử lý trách nhiệm người đứng đầu và cá nhân khác có trách nhiệm trong tổ chức chính trị, tổ chức chính trị - xã hội, tổ chức xã hội - nghề nghiệp và các tổ chức khác có sử dụng ngân sách nhà nước về việc để xảy ra hành vi tham nhũng được thực hiện theo quy định của Luật này và điều lệ, quy chế của tổ chức đó.

5. Trách nhiệm của người đứng đầu, cấp phó của người đứng đầu cơ quan, tổ chức, đơn vị quy định tại các khoản 1, 2 và 3 Điều này được loại trừ trong trường hợp họ không thể biết được hoặc đã áp dụng các biện pháp cần thiết để phòng ngừa, ngăn chặn hành vi tham nhũng.

Điều 55. Xử lý trách nhiệm người đứng đầu cơ quan, tổ chức, đơn vị khi để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách

1. Người đứng đầu cơ quan, tổ chức, đơn vị khi phải chịu trách nhiệm trực tiếp về việc để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách thì bị xử lý kỷ luật hoặc bị truy cứu trách nhiệm hình sự.

2. Người đứng đầu cơ quan, tổ chức, đơn vị khi phải chịu trách nhiệm liên đới về việc để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách thì bị xử lý kỷ luật.

3. Người đứng đầu cơ quan, tổ chức, đơn vị được xem xét miễn hoặc giảm trách nhiệm pháp lý quy định tại khoản 1 và khoản 2 Điều này nếu đã thực hiện các biện pháp cần thiết nhằm ngăn chặn, khắc phục hậu quả của hành vi tham nhũng; xử lý nghiêm minh, báo cáo kịp thời với cơ quan, tổ chức có thẩm quyền về hành vi tham nhũng.

4. Trong kết luận thanh tra, kết luận kiểm toán, kết luận điều tra vụ việc, vụ án tham nhũng phải nêu rõ trách nhiệm của người đứng đầu cơ quan, tổ chức, đơn vị để xảy ra hành vi tham nhũng theo các mức độ sau đây:

a) Yếu kém về năng lực quản lý;

b) Thiếu trách nhiệm trong quản lý;

c) Bao che cho người có hành vi tham nhũng.

Kết luận phải được gửi cho Ban chỉ đạo trung ương về phòng, chống tham nhũng, cơ quan, tổ chức có thẩm quyền.

Mục 6

CẢI CÁCH HÀNH CHÍNH, ĐỔI MỚI CÔNG NGHỆ QUẢN LÝ

VÀ PHƯƠNG THỨC THANH TOÁN

Điều 56. Cải cách hành chính nhằm phòng ngừa tham nhũng

Nhà nước thực hiện cải cách hành chính nhằm tăng cường tính độc lập và tự chịu trách nhiệm của cơ quan, tổ chức, đơn vị; đẩy mạnh việc phân cấp quản lý nhà nước giữa trung ương và địa phương, giữa các cấp chính quyền địa phương; phân định rõ nhiệm vụ, quyền hạn giữa các cơ quan nhà nước; công khai, đơn giản hoá và hoàn thiện thủ tục hành chính; quy định cụ thể trách nhiệm của từng chức danh trong cơ quan, tổ chức, đơn vị.

Điều 57. Tăng cường áp dụng khoa học, công nghệ trong quản lý

1. Cơ quan, tổ chức, đơn vị thường xuyên cải tiến công tác, tăng cường áp dụng khoa học, công nghệ trong hoạt động của mình, tạo thuận lợi để công dân, cơ quan, tổ chức, đơn vị thực hiện quyền và lợi ích hợp pháp của mình.

2. Cơ quan, tổ chức, đơn vị có trách nhiệm hướng dẫn trình tự, thủ tục giải quyết công việc để cơ quan, tổ chức, đơn vị, cá nhân chủ động thực hiện mà không phải trực tiếp tiếp xúc với cán bộ, công chức, viên chức.

Điều 58. Đổi mới phương thức thanh toán

1. Nhà nước áp dụng các biện pháp quản lý để thực hiện việc thanh toán thông qua tài khoản tại ngân hàng, Kho bạc nhà nước. Cơ quan, tổ chức, đơn vị có trách nhiệm thực hiện các quy định về thanh toán bằng chuyển khoản.

2. Chính phủ áp dụng các giải pháp tài chính, công nghệ tiến tới thực hiện mọi khoản chi đối với người có chức vụ, quyền hạn quy định tại các điểm a, b và c khoản 3 Điều 1 của Luật này và các giao dịch khác sử dụng ngân sách nhà nước phải thông qua tài khoản.

Chương III

PHÁT HIỆN THAM NHŨNG

Mục 1

CÔNG TÁC KIỂM TRA CỦA CƠ QUAN, TỔ CHỨC, ĐƠN VỊ

Điều 59. Công tác kiểm tra của cơ quan quản lý nhà nước

1. Thủ trưởng cơ quan quản lý nhà nước có trách nhiệm thường xuyên tổ chức kiểm tra việc chấp hành pháp luật của cơ quan, tổ chức, đơn vị, cá nhân thuộc phạm vi quản lý của mình nhằm kịp thời phát hiện hành vi tham nhũng.

2. Khi phát hiện có hành vi tham nhũng, thủ trưởng cơ quan quản lý nhà nước phải kịp thời xử lý theo thẩm quyền hoặc thông báo cho cơ quan thanh tra, điều tra, Viện kiểm sát có thẩm quyền.

Điều 60. Công tác tự kiểm tra của cơ quan, tổ chức, đơn vị

1. Người đứng đầu cơ quan, tổ chức, đơn vị có trách nhiệm chủ động tổ chức kiểm tra việc thực hiện nhiệm vụ, công vụ của cán bộ, công chức, viên chức thường xuyên, trực tiếp giải quyết công việc của cơ quan, tổ chức, đơn vị, cá nhân và cán bộ, công chức, viên chức khác do mình quản lý nhằm kịp thời phát hiện, ngăn chặn, xử lý hành vi tham nhũng.

2. Người đứng đầu cơ quan, tổ chức, đơn vị có trách nhiệm thường xuyên đôn đốc người đứng đầu đơn vị trực thuộc kiểm tra việc thực hiện nhiệm vụ, công vụ của cán bộ, công chức, viên chức do mình quản lý.

3. Khi phát hiện hành vi tham nhũng, người đứng đầu cơ quan, tổ chức, đơn vị phải kịp thời xử lý theo thẩm quyền hoặc thông báo cho cơ quan thanh tra, điều tra, Viện kiểm sát có thẩm quyền.

Điều 61. Hình thức kiểm tra

1. Việc kiểm tra thường xuyên được tiến hành theo chương trình, kế hoạch, tập trung vào lĩnh vực, hoạt động thường phát sinh hành vi tham nhũng.

2. Việc kiểm tra đột xuất được tiến hành khi phát hiện có dấu hiệu tham nhũng.

Mục 2

PHÁT HIỆN THAM NHŨNG THÔNG QUA HOẠT ĐỘNG THANH TRA,

KIỂM TOÁN, ĐIỀU TRA, KIỂM SÁT, XÉT XỬ, GIÁM SÁT

Điều 62. Phát hiện tham nhũng thông qua hoạt động thanh tra, kiểm toán, điều tra, kiểm sát, xét xử

Cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án thông qua hoạt động thanh tra, kiểm toán, điều tra, kiểm sát, xét xử có trách nhiệm chủ động phát hiện hành vi tham nhũng, xử lý theo thẩm quyền hoặc kiến nghị việc xử lý theo quy định của pháp luật và chịu trách nhiệm trước pháp luật về quyết định của mình.

Điều 63. Phát hiện tham nhũng thông qua hoạt động giám sát

Quốc hội, các cơ quan của Quốc hội, Đoàn đại biểu Quốc hội, Hội đồng nhân dân, đại biểu Quốc hội, đại biểu Hội đồng nhân dân thông qua hoạt động giám sát có trách nhiệm phát hiện hành vi tham nhũng, yêu cầu hoặc kiến nghị việc xử lý theo quy định của pháp luật.

Mục 3

TỐ CÁO VÀ GIẢI QUYẾT TỐ CÁO VỀ HÀNH VI THAM NHŨNG

Điều 64. Tố cáo hành vi tham nhũng và trách nhiệm của người tố cáo

1. Công dân có quyền tố cáo hành vi tham nhũng với cơ quan, tổ chức, cá nhân có thẩm quyền.

2. Người tố cáo phải tố cáo trung thực, nêu rõ họ, tên, địa chỉ, cung cấp thông tin, tài liệu mà mình có và hợp tác với cơ quan, tổ chức, cá nhân có thẩm quyền giải quyết tố cáo.

3. Người tố cáo mà cố tình tố cáo sai sự thật phải bị xử lý nghiêm minh, nếu gây thiệt hại cho người bị tố cáo thì phải bồi thường theo quy định của pháp luật.

Điều 65. Trách nhiệm tiếp nhận và giải quyết tố cáo

1. Cơ quan, tổ chức, đơn vị, cá nhân có trách nhiệm tạo điều kiện thuận lợi để công dân tố cáo trực tiếp, gửi đơn tố cáo, tố cáo qua điện thoại, tố cáo qua mạng thông tin điện tử và các hình thức khác theo quy định của pháp luật.

2. Người đứng đầu cơ quan, tổ chức có thẩm quyền khi nhận được tố cáo hành vi tham nhũng phải xem xét và xử lý theo thẩm quyền; giữ bí mật họ, tên, địa chỉ, bút tích và các thông tin khác theo yêu cầu của người tố cáo; áp dụng kịp thời các biện pháp cần thiết để bảo vệ người tố cáo khi có biểu hiện đe doạ, trả thù, trù dập người tố cáo hoặc khi người tố cáo yêu cầu; thông báo kết quả giải quyết tố cáo cho người tố cáo khi có yêu cầu.

3. Cơ quan thanh tra có trách nhiệm giúp thủ trưởng cơ quan quản lý nhà nước cùng cấp xác minh, kết luận về nội dung tố cáo và kiến nghị biện pháp xử lý; trong trường hợp phát hiện có dấu hiệu tội phạm thì chuyển cho cơ quan điều tra, Viện kiểm sát có thẩm quyền xử lý theo quy định của pháp luật về tố tụng hình sự.

Cơ quan điều tra, Viện kiểm sát nhận được tố cáo về hành vi tham nhũng phải xử lý theo thẩm quyền.

4. Thời hạn giải quyết tố cáo, thời hạn trả lời người tố cáo được thực hiện theo quy định của pháp luật.

Điều 66. Trách nhiệm phối hợp của cơ quan, tổ chức, đơn vị, cá nhân

Trong phạm vi nhiệm vụ, quyền hạn của mình, cơ quan, tổ chức, đơn vị, cá nhân phải tạo điều kiện, cộng tác với cơ quan, tổ chức, cá nhân có thẩm quyền giải quyết tố cáo để phát hiện, ngăn chặn và xử lý kịp thời hành vi tham nhũng, hạn chế thiệt hại do hành vi tham nhũng gây ra.

Điều 67. Khen thưởng người tố cáo

Người tố cáo trung thực, tích cực cộng tác với cơ quan, tổ chức, cá nhân có thẩm quyền trong việc phát hiện, ngăn chặn và xử lý hành vi tham nhũng thì được khen thưởng về vật chất, tinh thần theo quy định của pháp luật.

Chương IV

XỬ LÝ HÀNH VI THAM NHŨNG

VÀ CÁC HÀNH VI VI PHẠM PHÁP LUẬT KHÁC

Mục 1

XỬ LÝ KỶ LUẬT, XỬ LÝ HÌNH SỰ

Điều 68. Đối tượng bị xử lý kỷ luật, xử lý hình sự

1. Người có hành vi tham nhũng quy định tại Điều 3 của Luật này.

2. Người không báo cáo, tố giác khi biết được hành vi tham nhũng.

3. Người không xử lý báo cáo, tố giác, tố cáo về hành vi tham nhũng.

4. Người có hành vi đe doạ, trả thù, trù dập người phát hiện, báo cáo, tố giác, tố cáo, cung cấp thông tin về hành vi tham nhũng.

5. Người đứng đầu cơ quan, tổ chức, đơn vị để xảy ra hành vi tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý, phụ trách.

6. Người thực hiện hành vi khác vi phạm quy định của Luật này và quy định khác của pháp luật có liên quan.

Điều 69. Xử lý đối với người có hành vi tham nhũng

Người có hành vi tham nhũng thì tuỳ theo tính chất, mức độ vi phạm mà bị xử lý kỷ luật, truy cứu trách nhiệm hình sự; trong trường hợp bị kết án về hành vi tham nhũng và bản án, quyết định đã có hiệu lực pháp luật thì phải bị buộc thôi việc; đối với đại biểu Quốc hội, đại biểu Hội đồng nhân dân thì đương nhiên mất quyền đại biểu Quốc hội, đại biểu Hội đồng nhân dân.

Mục 2

XỬ LÝ TÀI SẢN THAM NHŨNG

Điều 70. Nguyên tắc xử lý tài sản tham nhũng

1. Cơ quan, tổ chức có thẩm quyền phải áp dụng các biện pháp cần thiết để thu hồi, tịch thu tài sản tham nhũng.

2. Tài sản tham nhũng phải được trả lại cho chủ sở hữu, người quản lý hợp pháp hoặc sung quỹ nhà nước.

3. Người đưa hối lộ mà chủ động khai báo trước khi bị phát hiện hành vi đưa hối lộ thì được trả lại tài sản đã dùng để hối lộ.

4. Việc tịch thu tài sản tham nhũng, thu hồi tài sản tham nhũng được thực hiện bằng quyết định của cơ quan nhà nước có thẩm quyền theo quy định của pháp luật.

Điều 71. Thu hồi tài sản tham nhũng có yếu tố nước ngoài

Trên cơ sở điều ước quốc tế mà Cộng hoà xã hội chủ nghĩa Việt Nam là thành viên và phù hợp với các nguyên tắc cơ bản của pháp luật Việt Nam, Chính phủ Việt Nam hợp tác với Chính phủ nước ngoài trong việc thu hồi tài sản của Việt Nam hoặc của nước ngoài bị tham nhũng và trả lại tài sản đó cho chủ sở hữu hợp pháp.

Chương V

TỔ CHỨC, TRÁCH NHIỆM VÀ HOẠT ĐỘNG PHỐI HỢP

CỦA CÁC CƠ QUAN THANH TRA, KIỂM TOÁN NHÀ NƯỚC,

ĐIỀU TRA, VIỆN KIỂM SÁT, TOÀ ÁN VÀ CỦA CƠ QUAN, TỔ CHỨC,

ĐƠN VỊ HỮU QUAN TRONG PHÒNG, CHỐNG THAM NHŨNG

Mục 1

TỔ CHỨC, CHỈ ĐẠO, PHỐI HỢP VÀ TRÁCH NHIỆM

TRONG CÔNG TÁC PHÒNG, CHỐNG THAM NHŨNG

Điều 72. Trách nhiệm của người đứng đầu cơ quan, tổ chức, đơn vị trong công tác phòng, chống tham nhũng

1. Người đứng đầu cơ quan, tổ chức, đơn vị có trách nhiệm áp dụng quy định của Luật này và các quy định khác của pháp luật có liên quan để tổ chức phòng, chống tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý.

2. Người đứng đầu cơ quan, tổ chức, đơn vị chịu trách nhiệm trước cơ quan, tổ chức, đơn vị cấp trên trực tiếp về việc phòng, chống tham nhũng trong cơ quan, tổ chức, đơn vị do mình quản lý.

Điều 73. Ban chỉ đạo phòng, chống tham nhũng

1. Ban chỉ đạo trung ương về phòng, chống tham nhũng do Thủ tướng Chính phủ đứng đầu có trách nhiệm chỉ đạo, phối hợp, kiểm tra, đôn đốc hoạt động phòng, chống tham nhũng trong phạm vi cả nước. Giúp việc cho Ban chỉ đạo trung ương về phòng, chống tham nhũng có bộ phận thường trực hoạt động chuyên trách.

2. Tổ chức, nhiệm vụ, quyền hạn và quy chế hoạt động của Ban chỉ đạo trung ương về phòng, chống tham nhũng do Uỷ ban thường vụ Quốc hội quy định theo đề nghị của Thủ tướng Chính phủ.

Điều 74. Giám sát công tác phòng, chống tham nhũng

1. Quốc hội, Uỷ ban thường vụ Quốc hội giám sát công tác phòng, chống tham nhũng trong phạm vi cả nước.

2. Hội đồng dân tộc và các Uỷ ban của Quốc hội trong phạm vi nhiệm vụ, quyền hạn của mình giám sát công tác phòng ngừa tham nhũng thuộc lĩnh vực do mình phụ trách.

Uỷ ban pháp luật của Quốc hội trong phạm vi nhiệm vụ, quyền hạn của mình giám sát việc phát hiện và xử lý hành vi tham nhũng.

3. Hội đồng nhân dân các cấp trong phạm vi nhiệm vụ, quyền hạn của mình có trách nhiệm giám sát công tác phòng, chống tham nhũng tại địa phương.

4. Đoàn đại biểu Quốc hội, đại biểu Quốc hội, đại biểu Hội đồng nhân dân trong phạm vi nhiệm vụ, quyền hạn của mình giám sát việc thực hiện các quy định của pháp luật về phòng, chống tham nhũng.

Điều 75. Đơn vị chuyên trách về chống tham nhũng

1. Trong Thanh tra Chính phủ, Bộ Công an, Viện kiểm sát nhân dân tối cao có đơn vị chuyên trách về chống tham nhũng.

2. Tổ chức, nhiệm vụ, quyền hạn của đơn vị chuyên trách về chống tham nhũng quy định tại khoản 1 Điều này do Uỷ ban thường vụ Quốc hội, Chính phủ quy định.

Điều 76. Trách nhiệm của Thanh tra Chính phủ

Trong phạm vi nhiệm vụ, quyền hạn của mình, Thanh tra Chính phủ có trách nhiệm sau đây:

1. Tổ chức, chỉ đạo, hướng dẫn công tác thanh tra việc thực hiện các quy định của pháp luật về phòng, chống tham nhũng; trường hợp phát hiện hành vi tham nhũng thì đề nghị cơ quan, tổ chức có thẩm quyền xử lý;

2. Xây dựng hệ thống dữ liệu chung về phòng, chống tham nhũng.

Điều 77. Trách nhiệm của Kiểm toán nhà nước

Trong phạm vi nhiệm vụ, quyền hạn của mình, Kiểm toán nhà nước có trách nhiệm tổ chức thực hiện việc kiểm toán nhằm phòng ngừa, phát hiện tham nhũng; trường hợp phát hiện hành vi tham nhũng thì đề nghị cơ quan, tổ chức có thẩm quyền xử lý.

Điều 78. Trách nhiệm của Bộ Công an, Bộ Quốc phòng

Trong phạm vi nhiệm vụ, quyền hạn của mình, Bộ Công an, Bộ Quốc phòng có trách nhiệm tổ chức, chỉ đạo thực hiện hoạt động điều tra tội phạm về tham nhũng.

Điều 79. Trách nhiệm của Viện kiểm sát nhân dân tối cao, Toà án nhân dân tối cao

1. Viện kiểm sát nhân dân tối cao có trách nhiệm tổ chức, chỉ đạo thực hiện hoạt động truy tố các tội phạm về tham nhũng; kiểm sát hoạt động điều tra, xét xử, thi hành án đối với các tội phạm về tham nhũng.

2. Toà án nhân dân tối cao có trách nhiệm xét xử, hướng dẫn công tác xét xử các tội phạm về tham nhũng.

Điều 80. Phối hợp hoạt động giữa các cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án

Cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án có trách nhiệm phối hợp trong phòng, chống tham nhũng theo các nội dung sau đây:

1. Trao đổi thường xuyên thông tin, tài liệu, kinh nghiệm về công tác phòng, chống tham nhũng;

2. Chuyển hồ sơ vụ việc tham nhũng cho cơ quan nhà nước có thẩm quyền xử lý;

3. Tổng hợp, đánh giá, dự báo tình hình tham nhũng và kiến nghị chính sách, giải pháp phòng, chống tham nhũng.

Điều 81. Phối hợp công tác giữa cơ quan thanh tra, kiểm toán nhà nước với cơ quan điều tra

1. Trong trường hợp cơ quan thanh tra, kiểm toán nhà nước chuyển hồ sơ vụ việc tham nhũng cho cơ quan điều tra thì cơ quan điều tra phải tiếp nhận và giải quyết theo quy định của pháp luật về tố tụng hình sự.

2. Trong trường hợp không đồng ý với việc giải quyết của cơ quan điều tra thì cơ quan thanh tra, kiểm toán nhà nước có quyền thông báo với Viện kiểm sát cùng cấp, cơ quan điều tra cấp trên.

Điều 82. Phối hợp công tác giữa cơ quan thanh tra, kiểm toán nhà nước với Viện kiểm sát

1. Trong trường hợp chuyển hồ sơ vụ việc tham nhũng cho cơ quan điều tra thì cơ quan thanh tra, kiểm toán nhà nước có trách nhiệm thông báo cho Viện kiểm sát cùng cấp để thực hiện việc kiểm sát.

2. Trong trường hợp cơ quan thanh tra, kiểm toán nhà nước chuyển hồ sơ vụ việc tham nhũng cho Viện kiểm sát thì Viện kiểm sát phải xem xét, giải quyết và thông báo kết quả giải quyết bằng văn bản cho cơ quan đã chuyển hồ sơ.

Mục 2

KIỂM TRA HOẠT ĐỘNG CHỐNG THAM NHŨNG

TRONG CƠ QUAN THANH TRA, KIỂM TOÁN NHÀ NƯỚC, ĐIỀU TRA,

VIỆN KIỂM SÁT, TOÀ ÁN

Điều 83. Kiểm tra hoạt động chống tham nhũng đối với cán bộ, công chức, viên chức của cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án

1. Cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án phải có biện pháp để kiểm tra nhằm ngăn chặn hành vi lạm quyền, lộng quyền, nhũng nhiễu của cán bộ, công chức, viên chức của mình trong hoạt động chống tham nhũng.

2. Người đứng đầu cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án phải tăng cường quản lý cán bộ, công chức, viên chức; chỉ đạo công tác thanh tra, kiểm tra nội bộ nhằm ngăn chặn hành vi vi phạm pháp luật trong hoạt động chống tham nhũng.

3. Cán bộ, công chức, viên chức của cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án có hành vi vi phạm pháp luật trong hoạt động chống tham nhũng thì tuỳ theo tính chất, mức độ vi phạm mà bị xử lý kỷ luật, truy cứu trách nhiệm hình sự; nếu gây thiệt hại thì phải bồi thường, bồi hoàn theo quy định của pháp luật.

Điều 84. Giải quyết tố cáo đối với cán bộ, công chức, viên chức của cơ quan thanh tra, kiểm toán nhà nước, điều tra, Vi��n kiểm sát, Toà án

Trường hợp có tố cáo về hành vi vi phạm pháp luật trong hoạt động chống tham nhũng đối với Thanh tra viên, Kiểm toán viên, Điều tra viên, Kiểm sát viên, Thẩm phán, Hội thẩm, Thư ký Toà án và cán bộ, công chức, viên chức khác của cơ quan thanh tra, kiểm toán nhà nước, điều tra, Viện kiểm sát, Toà án thì người đứng đầu cơ quan phải giải quyết theo thẩm quyền hoặc đề nghị cơ quan, tổ chức, cá nhân có thẩm quyền giải quyết.

Kết quả giải quyết tố cáo phải được công khai.

Chương VI

VAI TRÒ VÀ TRÁCH NHIỆM CỦA XÃ HỘI

TRONG PHÒNG, CHỐNG THAM NHŨNG

Điều 85. Vai trò và trách nhiệm của Mặt trận Tổ quốc Việt Nam và các tổ chức thành viên

1. Mặt trận Tổ quốc Việt Nam và các tổ chức thành viên có trách nhiệm sau đây:

a) Phối hợp với cơ quan nhà nước có thẩm quyền tuyên truyền, giáo dục nhân dân và các thành viên tổ chức mình thực hiện các quy định của pháp luật về phòng, chống tham nhũng; kiến nghị các biện pháp nhằm phát hiện và phòng ngừa tham nhũng;

b) Động viên nhân dân tham gia tích cực vào việc phát hiện, tố cáo hành vi tham nhũng;

c) Cung cấp thông tin và phối hợp với cơ quan, tổ chức, cá nhân có thẩm quyền trong việc phát hiện, xác minh, xử lý vụ việc tham nhũng;

d) Giám sát việc thực hiện pháp luật về phòng, chống tham nhũng.

2. Mặt trận Tổ quốc Việt Nam và các tổ chức thành viên có quyền yêu cầu cơ quan, tổ chức, cá nhân có thẩm quyền áp dụng biện pháp phòng ngừa tham nhũng, xác minh vụ việc tham nhũng, xử lý người có hành vi tham nhũng; cơ quan, tổ chức, cá nhân có thẩm quyền phải xem xét, trả lời trong thời hạn mười lăm ngày, kể từ ngày nhận được yêu cầu; trường hợp vụ việc phức tạp thì thời hạn trên có thể kéo dài nhưng không quá ba mươi ngày.

Điều 86. Vai trò và trách nhiệm của báo chí

1. Nhà nước khuyến khích cơ quan báo chí, phóng viên đưa tin phản ánh về vụ việc tham nhũng và hoạt động phòng, chống tham nhũng.

2. Cơ quan báo chí có trách nhiệm biểu dương tinh thần và những việc làm tích cực trong công tác phòng, chống tham nhũng; lên án, đấu tranh đối với những người có hành vi tham nhũng; tham gia tuyên truyền, phổ biến pháp luật về phòng, chống tham nhũng.

3. Cơ quan báo chí, phóng viên có quyền yêu cầu cơ quan, tổ chức, cá nhân có thẩm quyền cung cấp thông tin, tài liệu liên quan đến hành vi tham nhũng. Cơ quan, tổ chức, cá nhân được yêu cầu có trách nhiệm cung cấp thông tin, tài liệu đó theo quy định của pháp luật; trường hợp không cung cấp thì phải trả lời bằng văn bản và nêu rõ lý do.

4. Cơ quan báo chí, phóng viên phải đưa tin trung thực, khách quan. Tổng biên tập, phóng viên chịu trách nhiệm về việc đưa tin và chấp hành pháp luật về báo chí, quy tắc đạo đức nghề nghiệp.

Điều 87. Vai trò và trách nhiệm của doanh nghiệp, hiệp hội ngành nghề

1. Doanh nghiệp có trách nhiệm thông báo về hành vi tham nhũng và phối hợp với cơ quan, tổ chức, cá nhân có thẩm quyền trong việc xác minh, kết luận về hành vi tham nhũng.

2. Hiệp hội doanh nghiệp, hiệp hội ngành nghề có trách nhiệm tổ chức, động viên, khuyến khích hội viên của mình xây dựng văn hoá kinh doanh lành mạnh, phi tham nhũng.

3. Hiệp hội doanh nghiệp, hiệp hội ngành nghề và hội viên có trách nhiệm kiến nghị với Nhà nước hoàn thiện cơ chế, chính sách quản lý nhằm phòng, chống tham nhũng.

4. Nhà nước khuyến khích các doanh nghiệp cạnh tranh lành mạnh, có cơ chế kiểm soát nội bộ nhằm ngăn chặn hành vi tham ô, đưa hối lộ.

5. Cơ quan, tổ chức, cá nhân có thẩm quyền có trách nhiệm phối hợp với Phòng Thương mại và Công nghiệp Việt Nam, hiệp hội doanh nghiệp, hiệp hội ngành nghề và các tổ chức khác tổ chức diễn đàn để trao đổi, cung cấp thông tin, phục vụ công tác phòng, chống tham nhũng.

Điều 88. Trách nhiệm công dân, Ban thanh tra nhân dân

1. Công dân tự mình, thông qua Ban thanh tra nhân dân hoặc thông qua tổ chức mà mình là thành viên tham gia phòng, chống tham nhũng.

2. Ban thanh tra nhân dân tại xã, phường, thị trấn, trong cơ quan nhà nước, đơn vị sự nghiệp, doanh nghiệp của Nhà nước trong phạm vi nhiệm vụ, quyền hạn của mình giám sát việc thực hiện các quy định của pháp luật về phòng, chống tham nhũng.

Chương VII

HỢP TÁC QUỐC TẾ VỀ PHÒNG, CHỐNG THAM NHŨNG

Điều 89. Nguyên tắc chung về hợp tác quốc tế

Nhà nước cam kết thực hiện điều ước quốc tế về phòng, chống tham nhũng mà Cộng hoà xã hội chủ nghĩa Việt Nam là thành viên; hợp tác với các nước, tổ chức quốc tế, tổ chức, cá nhân nước ngoài trong hoạt động phòng, chống tham nhũng trên nguyên tắc tôn trọng độc lập, chủ quyền, toàn vẹn lãnh thổ và các bên cùng có lợi.

Điều 90. Trách nhiệm thực hiện hợp tác quốc tế

1. Thanh tra Chính phủ phối hợp với Bộ Ngoại giao, Bộ Công an và các cơ quan hữu quan thực hiện hoạt động hợp tác quốc tế về nghiên cứu, đào tạo, xây dựng chính sách, trao đổi thông tin, hỗ trợ tài chính, trợ giúp kỹ thuật, trao đổi kinh nghiệm trong phòng, chống tham nhũng.

2. Viện kiểm sát nhân dân tối cao, Bộ Tư pháp, Bộ Công an trong phạm vi nhiệm vụ, quyền hạn của mình thực hiện nhiệm vụ hợp tác quốc tế về tương trợ tư pháp trong phòng, chống tham nhũng.

Chương VIII

ĐIỀU KHOẢN THI HÀNH

Điều 91. Hiệu lực thi hành

1. Luật này có hiệu lực thi hành từ ngày 01 tháng 6 năm 2006.

2. Pháp lệnh chống tham nhũng ngày 26 tháng 02 năm 1998 và Pháp lệnh sửa đổi, bổ sung một số điều của Pháp lệnh chống tham nhũng ngày 28 tháng 4 năm 2000 hết hiệu lực kể từ ngày Luật này có hiệu lực.

Điều 92. Hướng dẫn thi hành

Chính phủ quy định chi tiết và hướng dẫn thi hành Luật này./.

Luật này đã được Quốc hội nước Cộng hoà xã hội chủ nghĩa Việt Nam khoá XI, kỳ họp thứ 8 thông qua ngày 29 tháng 11 năm 2005.
  	

CHỦ TỊCH QUỐC HỘI

(Đã ký)

Nguyễn Văn An`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}
