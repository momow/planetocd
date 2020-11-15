class CustomRender < Redcarpet::Render::HTML
    def link(link, title, content)
        "<a href=\"#{link}\" target=\"_blank\">#{content}</a>"
    end
end
