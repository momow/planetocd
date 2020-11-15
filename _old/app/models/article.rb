require 'redcarpet'
require_relative '../helpers/custom_render'

class Article
    attr_accessor :id, :language, :likes_count,
                    :original_url, :original_url_domain, :original_title, :original_author,
                    :original_url_host, :original_url_scheme_and_host,
                    :title, :slug, :public_notes_html, :content_html, :content_html_short,
                    :published
    
    @@markdown_html = Redcarpet::Markdown.new(CustomRender)

    def initialize(article_path)
        mdocd_content = IO.read(article_path)
        populate_from_mocd(mdocd_content)
        convert_to_html()
    end

    def likes_count
      Like.where(article_id: id).size
    end

    def populate_from_mocd(mdocd_content)
        @published = true
        parts = mdocd_content.split('===')
        @content_md = parts[1]
        lines = parts[0].split("\n")
        lines.each do |line|
            if line.start_with?('Id: ')
                @id = line[4..-1].to_i
            elsif line.start_with?('Published: ')
                @published = line[11..-1].strip.downcase != "false"
            elsif line.start_with?('Language: ')
                @language = line[10..-1]
            elsif line.start_with?('Original URL: ')
                @original_url = line[14..-1]
            elsif line.start_with?('Original title: ')
                @original_title = line[16..-1]
            elsif line.start_with?('Original author: ')
                @original_author = line[17..-1]
            elsif line.start_with?('Title: ')
                @title = line[7..-1]
            elsif line.start_with?('Translator notes: ')
                @public_notes_md = line[18..-1]
            end
        end

        uri = URI.parse(@original_url)
        @original_url_host = uri.host
        @original_url_scheme_and_host = "#{uri.scheme}://#{uri.host}"
        @slug = @title.parameterize
    end

    def convert_to_html
        char_idx = 500
        while char_idx + 1 < @content_md.length
            if @content_md[char_idx..char_idx+1] == "\n\n"
                break
            end
            char_idx += 1
        end
        char_idx += 1
        content_md_short = @content_md[0..char_idx]
        @content_html = @@markdown_html.render(@content_md).html_safe
        @content_html_short = @@markdown_html.render(content_md_short).html_safe
        if (@public_notes_md)
            @public_notes_html = @@markdown_html.render(@public_notes_md).html_safe
        end
    end

    def to_s
        "Id: #{@id}\n"\
        "Language: #{@language}\n"\
        "Published: #{@published}\n"\
        "Original Url: #{@original_url}\n"\
        "Original Title: #{@original_title}\n"\
        "Original Author: #{@original_author}\n"\
        "Title: #{@title}\n"\
        "Slug: #{@slug}\n"\
        "Translator Notes: #{@public_notes_md}\n"\
        "Raw MD: #{@content_md.length} characters\n"\
        "HTML: #{@content_html.length} characters\n"\
        "HTML (short): #{@content_html_short.length} characters\n"
    end

    def to_param
        [@id, @slug].join("-")
    end
end
